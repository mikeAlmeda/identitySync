package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/mikeAlmeda/rosterbridge-backend/internal/handlers"
)

func RegisterRoutes(r *chi.Mux) {
	r.Get("/students", handlers.GetStudents)
}
