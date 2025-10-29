package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mikeAlmeda/rosterbridge-backend/internal/db"
	"github.com/mikeAlmeda/rosterbridge-backend/internal/models"
)

func GetStudents(w http.ResponseWriter, r *http.Request) {
	cursor, err := db.StudentCollection.Find(context.Background(), map[string]any{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var students []models.Student
	if err := cursor.All(context.Background(), &students); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(students)
}
