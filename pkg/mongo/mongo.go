package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoStore struct {
	uri    string
	dbName string
	client *mongo.Client
}

func InitMongoStore(uri, dbName string) *mongoStore {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	var result bson.M
	if err = client.Database(dbName).RunCommand(context.Background(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Connected to DB")

	return &mongoStore{
		uri:    uri,
		dbName: dbName,
		client: client,
	}
}

func (s *mongoStore) Disconnect(ctx context.Context) error {
	return s.client.Disconnect(ctx)
}
