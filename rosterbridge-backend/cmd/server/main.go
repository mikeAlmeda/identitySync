package main

import (
	"log"
	"net/http"
	"os"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"github.com/mikeAlmeda/rosterbridge-backend/internal/db"
	"github.com/mikeAlmeda/rosterbridge-backend/internal/router"
)

func main() {
	// Load environment variables from .env file
	_ = godotenv.Load()

	// Connect to MongoDB
	db.InitMongo()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Register routes
	router.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
