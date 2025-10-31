package models

import (
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Student represents a student document in MongoDB.
type Student struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	SourceID   string             `bson:"source_id" json:"source_id"`
	FirstName  string             `bson:"first_name" json:"first_name"`
	LastName   string             `bson:"last_name" json:"last_name"`
	GradeLevel string             `bson:"grade_level" json:"grade_level"`
	SchoolID   string             `bson:"school_id" json:"school_id"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
	ChangeHash string             `bson:"change_hash" json:"change_hash"`
}

// studentJSON is the outward JSON representation (with id as string).
type studentJSON struct {
	ID         string    `json:"id"`
	SourceID   string    `json:"source_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	GradeLevel string    `json:"grade_level"`
	SchoolID   string    `json:"school_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	ChangeHash string    `json:"change_hash"`
}

// MashalJSON customizes the JSON marshalling for Student.
func (s *Student) MarshalJSON() ([]byte, error) {
	var id string
	if !s.ID.IsZero() {
		id = s.ID.Hex()
	}
	sj := studentJSON{
		ID:         id,
		SourceID:   s.SourceID,
		FirstName:  s.FirstName,
		LastName:   s.LastName,
		GradeLevel: s.GradeLevel,
		SchoolID:   s.SchoolID,
		CreatedAt:  s.CreatedAt,
		UpdatedAt:  s.UpdatedAt,
		ChangeHash: s.ChangeHash,
	}
	return json.Marshal(sj)
}
