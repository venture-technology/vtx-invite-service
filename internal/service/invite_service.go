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

func (i *InviteService) ReadAllInvites(ctx context.Context, cnh *string) ([]models.Invite, error) {
	return i.inviterepository.ReadAllInvites(ctx, cnh)
}

func (i *InviteService) AcceptedInvite(ctx context.Context, invite_id *int) error {
	return i.inviterepository.AcceptedInvite(ctx, invite_id)
}

func (i *InviteService) DeclineInvite(ctx context.Context, invite_id *int) error {
	return i.inviterepository.DeclineInvite(ctx, invite_id)
}

// Request in AccountManager to verify if school have the driver like employee. If they are partners, Employee is true, otherwise false.
func (i *InviteService) IsEmployee(ctx context.Context, cnh *string) bool {

	// This is a mock, at now.
	return false

}

// Validating both as a school and as a driver exist.
func CheckInviteEntities(invite *models.Invite) error {

	conf := config.Get()

	resp, err := http.Get(fmt.Sprintf("%s/%s", conf.Environment.School, invite.School.Name))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Print(resp.StatusCode)
		return fmt.Errorf("school is different")
	}

	resp, err = http.Get(fmt.Sprintf("%s/%s", conf.Environment.Driver, invite.Driver.Name))
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
