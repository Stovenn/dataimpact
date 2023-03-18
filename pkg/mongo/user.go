package mongo

import (
	"context"

	"github.com/stovenn/dataimpact/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userStore struct {
	*mongo.Collection
}

func NewUserStore() *userStore {
	return &userStore{C.Database("dataimpact").Collection("users")}
}

func (us *userStore) Create(user *model.User) error {
	_, err := us.Collection.InsertOne(context.Background(), user)
	return err
}

func (us *userStore) FindOne(id string) (*model.User, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	result := us.Collection.FindOne(context.Background(), filter)
	if err := result.Err(); err != nil {
		return nil, err
	}

	var u model.User
	err := result.Decode(&u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
