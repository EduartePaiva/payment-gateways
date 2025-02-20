package main

import (
	"context"
	"log"

	"github.com/EduartePaiva/payment-gateways/cmd"
	"github.com/EduartePaiva/payment-gateways/pkg/env"
	"github.com/EduartePaiva/payment-gateways/storage"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create mongo mongoClient
	mongoClient, err := mongo.Connect(options.Client().
		ApplyURI(env.Config.MongoURI))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	db := storage.NewDatabase(mongoClient)

	// create redis client
	opt, err := redis.ParseURL(env.Config.RedisURI)
	if err != nil {
		panic(err)
	}
	rdb := storage.NewRedisLocker(redis.NewClient(opt))

	log.Fatal(cmd.RunServer(ctx, db, rdb))
}
