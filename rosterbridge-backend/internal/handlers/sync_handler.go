package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mikeAlmeda/rosterbridge-backend/internal/services"
	"github.com/mikeAlmeda/rosterbridge-backend/internal/store"
)

// SyncHandler handles POST /sync requests.
func SyncHandler(st store.StudentStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// Create sync service
		syncService := services.NewSyncService(st)

		// Run sync
		result, err := syncService.SyncStudents(ctx)
		if err != nil {
			http.Error(w, "sync failed: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Return results as JSON
		w.Header().Set("Content-Type", "application/json")
		data, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	}
}
