package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/vtx-invites/internal/middleware"
	"github.com/venture-technology/vtx-invites/internal/service"
	"github.com/venture-technology/vtx-invites/models"
)

type InviteController struct {
	inviteservice *service.InviteService
}

func NewInviteController(inviteservice *service.InviteService) *InviteController {
	return &InviteController{
		inviteservice: inviteservice,
	}
}

func (ct *InviteController) RegisterRoutes(router *gin.Engine) {

	api := router.Group("api/v1")

	api.POST("/invite", middleware.SchoolMiddleware(), ct.CreateInvite)               // criando um convite para o motorista
	api.GET("/invite/:id", middleware.DriverMiddleware(), ct.ReadInvite)              // verificar um convite de escola
	api.GET("/invite", middleware.DriverMiddleware(), ct.FindAllInvitesDriverAccount) // verificar todos os convites feitos por escolas
	api.PATCH("/invite/:id", middleware.DriverMiddleware(), ct.AcceptedInvite)        // aceitar um convite de escola
	api.DELETE("/invite/:id", middleware.DriverMiddleware(), ct.DeclineInvite)        // recusar um convite de escola

}

func (ct *InviteController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (ct *InviteController) CreateInvite(c *gin.Context) {

	var input models.Invite

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	employee, err := ct.inviteservice.IsEmployee(c, &input)

	if err != nil {
		log.Printf("error while verify employee: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "internal server error"})
		return
	}

	if employee {
		log.Printf("employee is true: %v", employee)
		c.JSON(http.StatusBadRequest, gin.H{"message": "internal server error"})
		return
	}

	err = ct.inviteservice.InviteDriver(c, &input)

	if err != nil {
		log.Printf("error while creating invite: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, &input)

}

func (ct *InviteController) ReadInvite(c *gin.Context) {

	inviteId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Printf("converter error str to int: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"": ""})
	}

	invite, err := ct.inviteservice.ReadInvite(c, &inviteId)

	if err != nil {
		log.Printf("error while found invite: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invite don't found"})
		return
	}

	c.JSON(http.StatusOK, invite)

}

func (ct *InviteController) FindAllInvitesDriverAccount(c *gin.Context) {

	var input models.Driver

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	invites, err := ct.inviteservice.FindAllInvitesDriverAccount(c, &input.CNH)

	if err != nil {
		log.Printf("invites don't found: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "internal server error"})
		return
	}

	c.JSON(http.StatusAccepted, invites)

}

func (ct *InviteController) AcceptedInvite(c *gin.Context) {

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		log.Printf("error while convert string to int: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error at convert to int"})
		return
	}

	// read invite to get data of invite and create a partners between school and driver
	invite, err := ct.inviteservice.ReadInvite(c, &id)

	if err != nil {
		log.Printf("error while reading invite: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error at reading invite"})
		return
	}

	err = ct.inviteservice.AcceptedInvite(c, &id)

	if err != nil {
		log.Printf("error while accepting invite: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error at accepting invite"})
		return
	}

	err = ct.inviteservice.CreatePartner(c, invite)

	if err != nil {
		log.Printf("error while creating partner: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error at creating partner"})
		return
	}

	c.JSON(http.StatusCreated, invite)

}

func (ct *InviteController) DeclineInvite(c *gin.Context) {

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		log.Printf("error while convert string to int: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error at convert to int"})
		return
	}

	err = ct.inviteservice.DeclineInvite(c, &id)

	if err != nil {
		log.Printf("error while deleting invite: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error at deleting invite"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("invite deleted w/ successfully: %d", &id)})

}
