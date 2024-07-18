package service

import (
	"bytes"
	"context"
	"encoding/json"
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

func (i *InviteService) FindAllInvitesDriverAccount(ctx context.Context, cnh *string) ([]models.Invite, error) {
	return i.inviterepository.FindAllInvitesDriverAccount(ctx, cnh)
}

func (i *InviteService) AcceptedInvite(ctx context.Context, invite *models.Invite) error {

	err := i.CreatePartner(ctx, invite)

	if err != nil {
		return err
	}

	return i.inviterepository.AcceptedInvite(ctx, &invite.ID)
}

func (i *InviteService) DeclineInvite(ctx context.Context, invite_id *int) error {
	return i.inviterepository.DeclineInvite(ctx, invite_id)
}

// Request in AccountManager to verify if school have the driver like employee. If they are partners, Employee is true, otherwise false.
func (i *InviteService) IsEmployee(ctx context.Context, invite *models.Invite) (bool, error) {

	conf := config.Get()

	resp, err := http.Get(fmt.Sprintf("%s/%s?school=%s", conf.Environment.AccountManager, invite.Driver.CNH, invite.School.CNPJ))

	if err != nil {
		log.Printf("request error: %s", err.Error())
		return false, err
	}

	if resp.StatusCode == 200 {
		log.Printf("they are employers: %d", resp.StatusCode)
		return true, nil
	}

	return false, nil

}

// create partner between school and driver, then driver accepted invite, sending request to account manager
func (i *InviteService) CreatePartner(ctx context.Context, invite *models.Invite) error {

	conf := config.Get()

	jsonInvite, err := json.Marshal(invite)

	if err != nil {
		return err
	}

	resp, err := http.Post(fmt.Sprintf("%s/partner", conf.Environment.AccountManager), "application/json", bytes.NewBuffer(jsonInvite))

	if err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		return fmt.Errorf("request error: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	return nil

}
