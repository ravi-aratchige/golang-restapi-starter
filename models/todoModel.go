package models

import "gorm.io/gorm"

// Define todo model for database interaction
type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}
