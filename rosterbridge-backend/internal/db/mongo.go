package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var StudentCollection *mongo.Collection

func InitMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGO_URI is not set")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}

	Client = client
	StudentCollection = client.Database("rosterbridge").Collection("students")

	log.Println("Connected to MongoDB")
}
