package main

import (
	"log"
	"net/http"
	"time"

	config1 "github.com/htanmo/todoapi/internal/api/v1/config"
	v1 "github.com/htanmo/todoapi/internal/api/v1/routes"
	"github.com/rs/cors"
)

func main() {
	r := http.NewServeMux()

	// Loading different versioned configs
	config1.Load()

	// Register different versioned routes
	v1.RegisterTodoRoutes(r)

	// Setup CORS
	handle := cors.Default().Handler(r)

	// Server
	srv := &http.Server{
		Addr:         "localhost:8080",
		Handler:      handle,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("starting server running on :8080")
	log.Fatal(srv.ListenAndServe())
}
