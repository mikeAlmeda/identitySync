package store

import (
	"context"
	"time"

	"github.com/mikeAlmeda/rosterbridge-backend/internal/db"
	"github.com/mikeAlmeda/rosterbridge-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoStudentStore struct{}

// NewMongoStudentStore returns a StudentStore backed by the global db.StudentCollection.
func NewMongoStudentStore() StudentStore {
	return &mongoStudentStore{}
}

func (m *mongoStudentStore) GetAll(ctx context.Context, filter interface{}) ([]models.Student, error) {
	cursor, err := db.StudentCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var students []models.Student
	if err := cursor.All(ctx, &students); err != nil {
		return nil, err
	}
	return students, nil
}

func (m *mongoStudentStore) GetByID(ctx context.Context, id primitive.ObjectID) (*models.Student, error) {
	var s models.Student
	if err := db.StudentCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&s); err != nil {
		return nil, err
	}
	return &s, nil
}

func (m *mongoStudentStore) Upsert(ctx context.Context, s *models.Student) error {
	now := time.Now().UTC()
	if s.CreatedAt.IsZero() {
		s.CreatedAt = now
	}
	s.UpdatedAt = now

	filter := bson.M{}
	if !s.ID.IsZero() {
		filter["_id"] = s.ID
	} else if s.SourceID != "" {
		filter["source_id"] = s.SourceID
	}

	opts := options.Update().SetUpsert(true)
	update := bson.M{"$set": s}

	_, err := db.StudentCollection.UpdateOne(ctx, filter, update, opts)
	return err
}

func (m *mongoStudentStore) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := db.StudentCollection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
