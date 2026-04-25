package mongo

import (
	"api-std/config"
	"context"
	"fmt"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoPoolDestroy() {
	// FIXME
	ctx := context.TODO()
	MongoPool.Disconnect(ctx)
}

func MongoPoolInit() {
	credential := options.Credential{
		Username: config.Env.MongoUser,
		Password: config.Env.MongoPass,
	}

	clientOptions := options.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s:%s/", config.Env.MongoHost, config.Env.MongoPort)).
		SetAuth(credential).
		SetMaxPoolSize(50).
		SetMinPoolSize(10).
		SetMaxConnIdleTime(5 * time.Minute)

	ctx := context.Background()
	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Info("MongoDB Pool has been successfully initialized")
	MongoPool = mongoClient
}

func MongoPoolPing() error {
	ctx := context.Background()
	err := MongoPool.Ping(ctx, nil)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	return nil
}

var MongoPool *mongo.Client = nil
