package store

import (
	"context"

	"github.com/mikeAlmeda/rosterbridge-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MockStudentStore is a mock implementation for testing.
type MockStudentStore struct {
	GetAllFunc  func(ctx context.Context, filter interface{}) ([]models.Student, error)
	GetByIDFunc func(ctx context.Context, id primitive.ObjectID) (*models.Student, error)
	UpsertFunc  func(ctx context.Context, s *models.Student) error
	DeleteFunc  func(ctx context.Context, id primitive.ObjectID) error
}

func (m *MockStudentStore) GetAll(ctx context.Context, filter interface{}) ([]models.Student, error) {
	if m.GetAllFunc != nil {
		return m.GetAllFunc(ctx, filter)
	}
	return []models.Student{}, nil
}

func (m *MockStudentStore) GetByID(ctx context.Context, id primitive.ObjectID) (*models.Student, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(ctx, id)
	}
	return nil, nil
}

func (m *MockStudentStore) Upsert(ctx context.Context, s *models.Student) error {
	if m.UpsertFunc != nil {
		return m.UpsertFunc(ctx, s)
	}
	return nil
}

func (m *MockStudentStore) Delete(ctx context.Context, id primitive.ObjectID) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(ctx, id)
	}
	return nil
}
