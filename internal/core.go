package internal

import (
	"context"

	"github.com/stovenn/dataimpact/internal/model"
)

type Store interface {
	Create(ctx context.Context, user *model.User) error
	FindOne(ctx context.Context, id string) (*model.User, error)
}
