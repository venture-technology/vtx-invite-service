package service

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/venture-technology/vtx-invites/config"
	"github.com/venture-technology/vtx-invites/internal/repository"
	"github.com/venture-technology/vtx-invites/models"
)

type InviteService struct {
	inviterepository repository.IInviteRepository
	kafkarepository  repository.IKafkaRepository
}

func NewInviteService(repo repository.IInviteRepository, kafkarepo repository.IKafkaRepository) *InviteService {
	return &InviteService{
		inviterepository: repo,
		kafkarepository:  kafkarepo,
	}
}

func (i *InviteService) InviteDriver(ctx context.Context, invite *models.Invite) error {
	return i.inviterepository.InviteDriver(ctx, invite)
}

func (i *InviteService) ReadInvite(ctx context.Context, invite_id *int) (*models.Invite, error) {
	return i.inviterepository.ReadInvite(ctx, invite_id)
}

// find all invites of driver account
func (i *InviteService) FindAllInvitesDriverAccount(ctx context.Context, cnh *string) ([]models.Invite, error) {
	return i.inviterepository.FindAllInvitesDriverAccount(ctx, cnh)
}

func (i *InviteService) AcceptedInvite(ctx context.Context, invite_id *int) error {
	return i.inviterepository.AcceptedInvite(ctx, invite_id)
}

func (i *InviteService) DeclineInvite(ctx context.Context, invite_id *int) error {
	return i.inviterepository.DeclineInvite(ctx, invite_id)
}

// Request in AccountManager to verify if school have the driver like employee. If they are partners, Employee is true, otherwise false.
func (i *InviteService) IsEmployee(ctx context.Context, invite *models.Invite) error {

	// Add interpolate, with accountmangerurl + driver + school

	conf := config.Get()

	resp, _ := http.Get(fmt.Sprintf("%v", conf.Environment.AccountManager))

	if resp.StatusCode != 200 {
		return fmt.Errorf("request error: %d", resp.StatusCode)
	}

	return nil

}

// create partner between school and driver, then driver accepted invite
func (i *InviteService) CreatePartner(ctx context.Context, invite *models.Invite) error {
	return nil
}

// Validating both as a school and as a driver exist.
func CheckInviteEntities(invite *models.Invite) error {

	conf := config.Get()

	// Add interpolate = schoolurl + schoolcnpj
	resp, err := http.Get(fmt.Sprintf("%s/%s", conf.Environment.School, invite.School.CNPJ))
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("request error in school: %d", resp.StatusCode)
	}

	// Add interpolate = schoolurl + schoolcnpj
	resp, err = http.Get(fmt.Sprintf("%s/%s", conf.Environment.Driver, invite.Driver.CNH))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Print(resp.StatusCode)
		return fmt.Errorf("driver is different")
	}

	return nil

}
