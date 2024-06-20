package repository

import (
	"context"
	"database/sql"

	"github.com/venture-technology/vtx-invites/models"
)

type IInviteRepository interface {
	CreateInvite(ctx context.Context, invite *models.Invite) error
	ReadInvite(ctx context.Context, invite_id *int) (*models.Invite, error)
	ReadAllInvites(ctx context.Context, cnh *string) ([]models.Invite, error)
	AcceptedInvite(ctx context.Context, invite_id *int) error
	DeclineInvite(ctx context.Context, invite_id *int) error
	IsEmployee(ctx context.Context, cnh *string) error
}

type InviteRepository struct {
	db *sql.DB
}

func NewInviteRepository(db *sql.DB) *InviteRepository {
	return &InviteRepository{
		db: db,
	}
}

func (i *InviteRepository) CreateInvite(ctx context.Context, invite *models.Invite) error {
	return nil
}

func (i *InviteRepository) ReadInvite(ctx context.Context, invite_id *int) (*models.Invite, error) {
	return nil, nil
}

func (i *InviteRepository) ReadAllInvites(ctx context.Context, cnh *string) ([]models.Invite, error) {
	return []models.Invite{}, nil
}

func (i *InviteRepository) AcceptedInvite(ctx context.Context, invite_id *int) error {
	return nil
}

func (i *InviteRepository) DeclineInvite(ctx context.Context, invite_id *int) error {
	return nil
}

func (i *InviteRepository) IsEmployee(ctx context.Context, cnh *string) error {
	return nil
}
