package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/htanmo/todoapi/config"
	"github.com/htanmo/todoapi/controllers"
	"github.com/rs/cors"
)

func main() {
	// Load configs
	config.Load()

	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("GET /api/todo", controllers.GetAllTodos)
	mux.HandleFunc("POST /api/todo", controllers.CreateNewTodo)
	mux.HandleFunc("GET /api/todo/{id}", controllers.GetTodoById)
	mux.HandleFunc("PUT /api/todo/{id}", controllers.EditTodo)
	mux.HandleFunc("DELETE /api/todo/{id}", controllers.DeleteTodo)

	// Setup CORS
	r := cors.Default().Handler(mux)

	// Get port number
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Server setup
	srv := &http.Server{
		Addr:         "localhost:" + port,
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("starting server running on :" + port)
	log.Fatal(srv.ListenAndServe())
}
