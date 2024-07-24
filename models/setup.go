package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(dbname string) {
	db, err := gorm.Open(sqlite.Open(dbname), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	err = db.AutoMigrate(&Todo{})
	if err != nil {
		return
	}
	DB = db
}
