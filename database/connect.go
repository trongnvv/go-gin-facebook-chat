package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Init() {
	MONGODB_URI := os.Getenv("MONGODB_URI")
	MONGODB_NAME := os.Getenv("MONGODB_NAME")

	client, err := mongo.NewClient(options.Client().ApplyURI(MONGODB_URI))
	if err != nil {
		log.Fatal(err)
	}
	//TODO
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongodb connected!")
	db := client.Database(MONGODB_NAME)
	db.Client().Disconnect(context.TODO())
	initModel(db)
}

func initModel(db *mongo.Database) {
	// models.UserModel = db.Collection("chat")
}
