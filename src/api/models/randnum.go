package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RandNum struct {
	ID      uint `json:"id" gorm:"primary_key"`
	RandNum int  `json:"randNum" gorm:"randNum"`
}

type MongoRandNum struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	RandNum   int                `json:"randNum" bson:"randNum"`
}
