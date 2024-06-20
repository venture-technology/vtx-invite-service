package service

import (
	"context"

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

func (i *InviteService) CreateInvite(ctx context.Context, invite *models.Invite) error {
	return i.inviterepository.CreateInvite(ctx, invite)
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

func (i *InviteService) IsEmployee(ctx context.Context, cnh *string) error {
	return i.inviterepository.IsEmployee(ctx, cnh)
}
