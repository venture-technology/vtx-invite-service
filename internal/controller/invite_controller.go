package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/venture-technology/vtx-invites/internal/middleware"
	"github.com/venture-technology/vtx-invites/internal/service"
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

	api.POST("/invite", middleware.SchoolMiddleware(), ct.CreateInvite)        // criando um convite para o motorista
	api.GET("/invite/:id", middleware.DriverMiddleware(), ct.ReadInvite)       // verificar todos os convites feitos por escolas
	api.GET("/invite", middleware.DriverMiddleware(), ct.ReadAllInvites)       // verificar um convite de escola
	api.PATCH("/invite/:id", middleware.DriverMiddleware(), ct.AcceptedInvite) // aceitar um convite de escola
	api.DELETE("/invite/:id", middleware.DriverMiddleware(), ct.DeclineInvite) // recusar um convite de escola

}

func (ct *InviteController) CreateInvite(c *gin.Context) {

}

func (ct *InviteController) ReadInvite(c *gin.Context) {

}

func (ct *InviteController) ReadAllInvites(c *gin.Context) {

}

func (ct *InviteController) AcceptedInvite(c *gin.Context) {

}

func (ct *InviteController) DeclineInvite(c *gin.Context) {

}
