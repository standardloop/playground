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

var mongoClient *mongo.Client
var collection *mongo.Collection
var ctx = context.TODO()

func Init() {
	credential := options.Credential{
		Username: config.Env.MongoUser,
		Password: config.Env.MongoPass,
	}
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s/", config.Env.MongoHost, config.Env.MongoPort)).SetAuth(credential)
	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal().Msg("cannot make client to mongodb")
	}
	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		log.Error().Msg("Failed to ping mongodb")
	} else {
		log.Debug().Msg("succsesfully pinged mongodb")
	}
	collection = mongoClient.Database(config.Env.MongoDBName).Collection("randNum")
	dbSeed()
}

func dbSeed() {
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

func HealthCheck() error {
	err := mongoClient.Ping(ctx, nil)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// this function is really only a POC
func GetOne() ([]*models.MongoRandNum, error) {
	pipeline := []bson.D{bson.D{{Key: "$sample", Value: bson.D{{Key: "size", Value: 1}}}}}
	var randNums []*models.MongoRandNum
	cur, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Error().Msg(err.Error())
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

// not needed right now, but save for later
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
