package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Connect() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	c, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = c.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	client = c
}

func GetCollection(collectionName string) *mongo.Collection {
	return client.Database("carDB").Collection(collectionName)
}
