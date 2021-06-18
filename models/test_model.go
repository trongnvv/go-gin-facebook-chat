package models

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var TestModel *mongo.Collection

type TestSchema struct {
	*BaseSchema `bson:",inline"`
	Name    string `bson:"name"`
}
