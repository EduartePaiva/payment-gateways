package main

import (
	"context"
	"log"

	"github.com/EduartePaiva/payment-gateways/cmd"
	"github.com/EduartePaiva/payment-gateways/pkg/env"
	"github.com/EduartePaiva/payment-gateways/storage"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := mongo.Connect(options.Client().
		ApplyURI(env.Config.MongoURI))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	db := storage.NewDatabase(client)

	log.Fatal(cmd.RunServer(ctx, db))
}
