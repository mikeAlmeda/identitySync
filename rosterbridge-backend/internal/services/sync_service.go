package services

import (
	"context"
	"log"
	"time"

	"github.com/mikeAlmeda/rosterbridge-backend/internal/models"
	"github.com/mikeAlmeda/rosterbridge-backend/internal/store"
)

// SyncResults contains statistics from a sync operation.
type SyncResults struct {
	Synced  int `json:"synced"`
	Created int `json:"created"`
	Updated int `json:"updated"`
	Errors  int `json:"errors"`
}

// SyncService handles roster synchronization logic.
type SyncService struct {
	store store.StudentStore
}

// NewSync Service creates a new sync service.
func NewSyncService(st store.StudentStore) *SyncService {
	return &SyncService{store: st}
}

// SyncStudents fetches external roster data and upserts to database
func (s *SyncService) SyncStudents(ctx context.Context) (*SyncResults, error) {
	// Mock external API data (simulates Clever or similar)
	externalStudents := s.fetchMockRosterData()

	result := &SyncResults{}

	for _, extStudent := range externalStudents {
		// Check if student exists by source_id
		existing, err := s.findBySourceID(ctx, extStudent.SourceID)
		if err != nil {
			log.Printf("Error checking student %s: %v", extStudent.SourceID, err)
			result.Errors++
			continue
		}

		// Transform external data to our Student model
		student := models.Student{
			SourceID:  extStudent.SourceID,
			FirstName: extStudent.FirstName,
			LastName:  extStudent.LastName,
			Email:     extStudent.Email,
			UpdatedAt: time.Now().UTC(),
		}

		if existing != nil {
			// Update existing student
			student.ID = existing.ID
			student.CreatedAt = existing.CreatedAt
			result.Updated++
		} else {
			// New student
			student.CreatedAt = time.Now().UTC()
			result.Created++
		}

		// Upsert
		if err := s.store.Upsert(ctx, &student); err != nil {
			log.Printf("Error upserting student %s: %v", student.SourceID, err)
			result.Errors++
			continue
		}

		result.Synced++
	}

	return result, nil
}

// findBySourceID checks if a student exists by source_id.
func (s *SyncService) findBySourceID(ctx context.Context, sourceID string) (*models.Student, error) {
	// Use GetAll with filter (MongoDB style)
	filter := map[string]interface{}{"source_id": sourceID}
	students, err := s.store.GetAll(ctx, filter)
	if err != nil {
		return nil, err
	}
	if len(students) > 0 {
		return &students[0], nil
	}
	return nil, nil
}

// fetchMockRosterData simulates fetching from an external API.
func (s *SyncService) fetchMockRosterData() []ExternalStudent {
	return []ExternalStudent{
		{SourceID: "clever-001", FirstName: "Alice", LastName: "Johnson", Grade: 9},
		{SourceID: "clever-002", FirstName: "Bob", LastName: "Smith", Grade: 10},
		{SourceID: "clever-003", FirstName: "Charlie", LastName: "Brown", Grade: 11},
		{SourceID: "clever-004", FirstName: "Homer", LastName: "Simpson", Grade: 12},
		{SourceID: "clever-005", FirstName: "Barney", LastName: "Gumble", Grade: 8},
	}
}

// ExternalStudent represents data forma n external roster API.
type ExternalStudent struct {
	SourceID  string
	FirstName string
	LastName  string
	Email     string
	Grade     int
}
