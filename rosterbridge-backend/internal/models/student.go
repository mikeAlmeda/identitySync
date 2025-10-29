package models

type Student struct {
	ID         string `json:"id" bson:"id"`
	FirstName  string `json:"first_name" bson:"first_name"`
	LastName   string `json:"last_name" bson:"last_name"`
	GradeLevel string `json:"grade_level" bson:"grade_level"`
	SchoolID   string `json:"school_id" bson:"school_id"`
}
