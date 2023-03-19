package internal

import (
	"context"

	"github.com/stovenn/dataimpact/internal/model"
)

type Store interface {
	Create(ctx context.Context, user *model.User) error
	FindOne(ctx context.Context, id string) (*model.User, error)
	Find(ctx context.Context) ([]*model.User, error)
	Update(ctx context.Context, id string, update *model.User) error
	DeleteOne(ctx context.Context, id string) error
}
