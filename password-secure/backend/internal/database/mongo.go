package config

import (
	"context"
	"log"
	"pwd-safe/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
			log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
}

func SavePassword(p models.Password) error {
	collection := client.Database("passwords_db").Collection("passwords")
	_, err := collection.InsertOne(context.Background(), p)
	return err
}
