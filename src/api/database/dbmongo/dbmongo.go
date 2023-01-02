package dbmongo

import (
	"api/config"
	"api/models"
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient = dbInit()
var collection *mongo.Collection
var ctx = context.TODO()

func DBSeed() {
	collection = MongoClient.Database(config.Env.MongoDBName).Collection("randNum")
	for i := 1; i < 100; i++ {
		randNum := &models.MongoRandNum{
			ID:        primitive.NewObjectID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			RandNum:   rand.Intn(100 - 0),
		}
		err := createRandNum(randNum)
		if err != nil {
			log.Error().Msg("rip mongo seed")
		}
	}
}

func createRandNum(randNum *models.MongoRandNum) error {
	_, err := collection.InsertOne(ctx, randNum)
	return err
}

func getAll() ([]*models.MongoRandNum, error) {
	filter := bson.D{{}}
	return filterRandNums(filter)
}
func filterRandNums(filter interface{}) ([]*models.MongoRandNum, error) {
	var randNums []*models.MongoRandNum
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return randNums, err
	}
	for cur.Next(ctx) {
		var t models.MongoRandNum
		err := cur.Decode(&t)
		if err != nil {
			return randNums, err
		}
		randNums = append(randNums, &t)
	}
	if err := cur.Err(); err != nil {
		return randNums, err
	}
	cur.Close(ctx)
	if len(randNums) == 0 {
		return randNums, mongo.ErrNoDocuments
	}
	return randNums, nil
}

func dbInit() *mongo.Client {
	credential := options.Credential{
		Username: config.Env.MongoUser,
		Password: config.Env.MongoPass,
	}
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s/", config.Env.MongoHost, config.Env.MongoPort)).SetAuth(credential)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal().Msg("cannot make client to mongodb")
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Error().Msg("Failed to ping mongodb")
	} else {
		log.Debug().Msg("succsesfully pinged mongodb")
	}
	return client
}
