package routes

import (
	"net/http"

	"github.com/htanmo/todoapi/internal/api/v1/controllers"
)

func RegisterTodoRoutes(mux *http.ServeMux) {
	// Register routes
	mux.HandleFunc("GET /api/v1/todo", controllers.GetAllTodos)
	mux.HandleFunc("POST /api/v1/todo", controllers.CreateNewTodo)
	mux.HandleFunc("GET /api/v1/todo/{id}", controllers.GetTodoById)
	mux.HandleFunc("PUT /api/v1/todo/{id}", controllers.EditTodo)
	mux.HandleFunc("DELETE /api/v1/todo/{id}", controllers.DeleteTodo)
}
