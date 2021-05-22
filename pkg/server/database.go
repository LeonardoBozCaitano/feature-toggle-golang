package server

import (
	"context"
	"log"

	"github.com/feature_toggle/pkg/config"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

func ConnectToDatabase() (*mongo.Database, error) {
	cfg := config.NewConfig()

	clientOptions := options.Client().ApplyURI(cfg.DatabaseURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(cfg.DatabaseName), err
}
