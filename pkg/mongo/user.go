package mongo

import (
	"context"
	"fmt"

	"github.com/stovenn/dataimpact/pkg/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore interface {
	Create()
}

type userStore struct {
	*mongo.Collection
}

func NewUserStore(client *mongo.Client) *userStore {
	collection := client.Database("dataimpact").Collection("users")
	return &userStore{collection}
}

func (us *userStore) Create() {
	fmt.Println("create")
	user := model.User{
		ID:   "sfdskh23hreef8d7fh",
		Name: "Sam Lee",
		Tags: []string{"seasons", "gardening", "flower"},
	}

	_, err := us.InsertOne(context.Background(), user)
	if err != nil {
		panic(err)
	}
}
