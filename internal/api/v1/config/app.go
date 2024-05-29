package config

import (
	"log"
	"os"

	"github.com/htanmo/todoapi/internal/api/v1/models"
	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file.")
	}

	dbURI := os.Getenv("DB_URI")

	// default db uri
	if dbURI == "" {
		dbURI = "todo.db"
	}

	models.ConnectDatabase(dbURI)
}
