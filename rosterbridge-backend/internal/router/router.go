package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/mikeAlmeda/rosterbridge-backend/internal/handlers"
	"github.com/mikeAlmeda/rosterbridge-backend/internal/store"
)

func RegisterRoutes(r chi.Router) {
	st := store.NewMongoStudentStore()

	r.Get("/students", handlers.GetStudentsHandler(st))

	// Additional routes can be registered here, injecting 'st' as needed.
}
