package mongo

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stovenn/dataimpact/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var testStore *mongoStore

func TestMain(m *testing.M) {
	config, err := util.SetupConfig("../..")
	if err != nil {
		log.Fatalf("cannot load config: %v\n", err)
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(config.DBUri).SetServerAPIOptions(serverAPI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	var result bson.M
	if err = client.Database(config.DBName).RunCommand(context.Background(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}

	testStore = &mongoStore{
		uri:    config.DBUri,
		dbName: config.DBName,
		client: client,
	}

	os.Exit(m.Run())
}
