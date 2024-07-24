package config

import (
	"os"

	"github.com/htanmo/todoapi/models"
	"github.com/joho/godotenv"
)

func Load() {
	godotenv.Load()

	dbURI := os.Getenv("DB_URI")

	// Default db uri
	if dbURI == "" {
		dbURI = "todo.db"
	}

	models.ConnectDatabase(dbURI)
}
