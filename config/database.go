package config

import (
	"rawdog-web-server/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToDatabase() gorm.DB {
	// Connect to the SQLite database
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	// Auto-migrate models to create tables
	db.AutoMigrate(&models.Todo{})

	return *db
}
