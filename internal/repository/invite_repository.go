package repository

import (
	"context"

	"github.com/venture-technology/vtx-invites/models"
)

type IDriverRepository interface {
	IsEmployee(ctx context.Context, cnh *string) error
	CreateInvite(ctx context.Context, invite *models.Invite) error
	ReadInvite(ctx context.Context, invite_id *int) (*models.Invite, error)
	ReadAllInvites(ctx context.Context, cnh *string) ([]models.Invite, error)
	UpdateInvite(ctx context.Context, invite_id *int) error
	DeleteInvite(ctx context.Context, invite_id *int) error
}
