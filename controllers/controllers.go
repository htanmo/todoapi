package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/htanmo/todoapi/models"
)

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var todos []models.Todo
	models.DB.Find(&todos)
	json.NewEncoder(w).Encode(todos)
}

func CreateNewTodo(w http.ResponseWriter, r *http.Request) {
	var newtodo models.Todo
	err := json.NewDecoder(r.Body).Decode(&newtodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.DB.Create(&newtodo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newtodo)
}

func GetTodoById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(r.PathValue("id"), 0, 0)
	var todo models.Todo
	if err := models.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		http.Error(w, "error: Record not found!", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&todo)
}

func EditTodo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(r.PathValue("id"), 0, 0)
	var todo models.Todo
	if err := models.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		http.Error(w, "error: Record not found!", http.StatusBadRequest)
		return
	}
	var editTodo models.Todo
	err := json.NewDecoder(r.Body).Decode(&editTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.DB.Model(&todo).Updates(editTodo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(r.PathValue("id"), 0, 0)
	var todo models.Todo
	if err := models.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		http.Error(w, "error: Record not found", http.StatusBadRequest)
		return
	}
	models.DB.Delete(&todo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	result := struct {
		Data bool `json:"data"`
	}{
		Data: true,
	}
	json.NewEncoder(w).Encode(result)
}
