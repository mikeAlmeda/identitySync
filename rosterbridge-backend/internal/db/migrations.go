package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// EnsureIndexes creates indexes if they don't exist (idempotent)
func EnsureIndexes(ctx context.Context) error {
	indexes := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "source_id", Value: 1}},
			Options: options.Index().SetUnique(true).SetName("idx_source_id"),
		},
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetName("idx_email"),
		},
		{
			Keys:    bson.D{{Key: "updated_at", Value: -1}},
			Options: options.Index().SetName("idx_updated_at_desc"),
		},
	}

	_, err := StudentCollection.Indexes().CreateMany(ctx, indexes)
	if err != nil {
		return err
	}
	log.Println("indexes ensured for students collection")
	return nil
}
