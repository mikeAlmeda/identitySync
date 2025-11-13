package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

var Client *mongo.Client
var StudentCollection *mongo.Collection

// Connect creates a Mongo client, pings, sets collections and returns error if any.
func Connect(ctx context.Context) (*mongo.Client, error) {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return nil, errMissingURI()
	}

	clientOpts := options.Client().ApplyURI(uri)
	// optional: add write concern or other options
	clientOpts.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))

	ctxConnect, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctxConnect, clientOpts)
	if err != nil {
		return nil, err
	}

	Client = client
	StudentCollection = client.Database("rosterbridge").Collection("students")
	log.Println("Connected to MongoDB")
	return client, nil
}

func Disconnect(ctx context.Context) error {
	if Client == nil {
		return nil
	}
	ctxTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	return Client.Disconnect(ctxTimeout)
}

func errMissingURI() error {
	return &MissingEnvError{"MONGODB_URI"}
}

// MissingEnvError represents an error for missing environment variables.
type MissingEnvError struct{ Var string }

func (e *MissingEnvError) Error() string { return e.Var + " is not set" }

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
