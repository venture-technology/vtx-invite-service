package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	// This is a mock, at moment.
	conf := config.Get()

	resp, err := http.Get(fmt.Sprintf("%s/driver/%s?school=%s", conf.Environment.AccountManager, invite.Driver.CNH, invite.School.CNPJ))

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var response models.Response
	if err := json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("erro ao decodificar o JSON: %v", err.Error())
	}

	err = processPayout(response.Payout)
	if err != nil {
		return fmt.Errorf("erro: %v", err)
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

	resp, err := http.Get(fmt.Sprintf("%s/%s", conf.Environment.AccountManager, invite.School.Name))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Print(resp.StatusCode)
		return fmt.Errorf("school is different")
	}

	resp, err = http.Get(fmt.Sprintf("%s/%s", conf.Environment.AccountManager, invite.Driver.Name))
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

func processPayout(payout *models.Payout) error {

	if payout == nil {
		return nil
	}

	if payout.Driver != nil && payout.School != nil {
		return fmt.Errorf("school and driver are partners")
	}

	return fmt.Errorf("error to check processPayout")

}
