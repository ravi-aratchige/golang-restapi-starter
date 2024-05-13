package config

import (
	"golang-restapi-starter/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDbConnection() gorm.DB {
	// Connect to the SQLite database
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto-migrate the Todo model to create the table
	db.AutoMigrate(&models.Todo{})

	return *db
}
