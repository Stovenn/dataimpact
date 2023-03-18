package internal

import "github.com/stovenn/dataimpact/internal/model"

type UserStore interface {
	Create(user *model.User) error
	FindOne(id string) (*model.User, error)
}
