package mongo

import (
	"context"
	"errors"

	"github.com/stovenn/dataimpact/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ms *mongoStore) Create(ctx context.Context, user *model.User) error {
	_, err := ms.userCollection().InsertOne(ctx, user)

	return err
}

func (ms *mongoStore) FindOne(ctx context.Context, id string) (*model.User, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	result := ms.userCollection().FindOne(ctx, filter)
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

func (ms *mongoStore) Find(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	cursor, err := ms.userCollection().Find(ctx, bson.D{{}})
	if err != nil {
		return users, nil
	}
	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ms *mongoStore) DeleteOne(ctx context.Context, id string) error {
	filter := bson.D{{Key: "_id", Value: id}}
	_, err := ms.userCollection().DeleteOne(ctx, filter)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	}

	return nil
}

func (ms *mongoStore) userCollection() *mongo.Collection {
	return ms.client.Database(ms.dbName).Collection("users")
}
