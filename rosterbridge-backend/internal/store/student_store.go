package store

import (
	"context"

	"github.com/mikeAlmeda/rosterbridge-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetStudentByID retrieves a student by their ObjectID.
type StudentStore interface {
	GetAll(ctx context.Context, filter interface{}) ([]models.Student, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (*models.Student, error)
	Upsert(ctx context.Context, s *models.Student) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}
