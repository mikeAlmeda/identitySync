package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/mikeAlmeda/rosterbridge-backend/internal/models"
	"github.com/mikeAlmeda/rosterbridge-backend/internal/store"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetStudentsHandler_Success(t *testing.T) {
	// Setup: mock store returns fake students
	mockStore := &store.MockStudentStore{
		GetAllFunc: func(ctx context.Context, filter interface{}) ([]models.Student, error) {
			return []models.Student{
				{
					ID:        primitive.NewObjectID(),
					SourceID:  "student-123",
					FirstName: "Alice",
					LastName:  "Smith",
					Email:     "alice@example.com",
					CreatedAt: time.Now().UTC(),
					UpdatedAt: time.Now().UTC(),
				},
				{
					ID:        primitive.NewObjectID(),
					SourceID:  "student-456",
					FirstName: "Bob",
					LastName:  "Johnson",
					Email:     "bob@example.com",
					CreatedAt: time.Now().UTC(),
					UpdatedAt: time.Now().UTC(),
				},
			}, nil
		},
	}

	// Create handler with mock store
	handler := GetStudentsHandler(mockStore)

	// Create test request
	req := httptest.NewRequest(http.MethodGet, "/students", nil)
	w := httptest.NewRecorder()

	// Execute handler
	handler.ServeHTTP(w, req)

	// Assert status code
	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	// Assert Content-Type
	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", contentType)
	}

	// Assert JSON response
	var students []models.Student
	if err := json.NewDecoder(w.Body).Decode(&students); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(students) != 2 {
		t.Errorf("expected 2 students, got %d", len(students))
	}

	if students[0].FirstName != "Alice" {
		t.Errorf("expected first student name Alice, got %s", students[0].FirstName)
	}
}

// Test for error case
func TestGetStudentsHandler_Error(t *testing.T) {
	// Setup: mock store returns an error
	mockStore := &store.MockStudentStore{
		GetAllFunc: func(ctx context.Context, filter interface{}) ([]models.Student, error) {
			return nil, errors.New("database connection failed")
		},
	}

	handler := GetStudentsHandler(mockStore)

	req := httptest.NewRequest(http.MethodGet, "/students", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	// Assert status code is 500
	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", w.Code)
	}

	// Assert error message in body
	body := w.Body.String()
	if body == "" {
		t.Error("expected error message in response body")
	}
}

// Test for empty result case
func TestGetStudentsHandler_EmptyResult(t *testing.T) {
	// Setup: mock store returns empty slice
	mockStore := &store.MockStudentStore{
		GetAllFunc: func(ctx context.Context, filter interface{}) ([]models.Student, error) {
			return []models.Student{}, nil
		},
	}

	handler := GetStudentsHandler(mockStore)

	req := httptest.NewRequest(http.MethodGet, "/students", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var students []models.Student
	if err := json.NewDecoder(w.Body).Decode(&students); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(students) != 0 {
		t.Errorf("expected 0 students, got %d", len(students))
	}
}
