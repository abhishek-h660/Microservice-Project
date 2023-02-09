package models

import (
	"Microservice/logger/utils"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Log struct {
	ID        string    `json:"_id,ommit_empty" bson:"_id,ommit_empty"`
	Name      string    `json:"name" bson:"name"`
	Data      string    `json:"data" bson:"data"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func CollectionLogs(db *mongo.Database) *mongo.Collection {
	return db.Collection("logs")
}

func (l *Log) Insert(db *mongo.Database) (*mongo.InsertOneResult, error) {
	return CollectionLogs(db).InsertOne(context.Background(), l)
}

func All(db *mongo.Database) []Log {
	var logs []Log
	cur, err := CollectionLogs(db).Aggregate(context.Background(), []bson.M{
		{
			"sort": bson.M{"created_at": -1},
		},
	})

	utils.HandleErrors(err)

	err = cur.All(context.Background(), &logs)
	utils.HandleErrors(err)

	return logs
}

func GetOne(db *mongo.Database, id string) Log {
	var log Log
	_id, err := primitive.ObjectIDFromHex(id)
	utils.HandleErrors(err)

	err = CollectionLogs(db).FindOne(context.Background(), bson.M{"_id": _id}).Decode(&log)
	utils.HandleErrors(err)
	return log
}
