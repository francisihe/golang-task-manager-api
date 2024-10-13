package models

import (
	"time"

	"gorm.io/gorm"
)

// Global variable to hold the database connection
var DB *gorm.DB

// Task represents a task model with GORM tags for database mapping
type Task struct {
	ID          string    `gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// package models

// import "time"

// // Task struct to define the task model
// type Task struct {
// 	ID          string    `json:"id"`
// 	Title       string    `json:"title"`
// 	Description string    `json:"description"`
// 	Status      string    `json:"status"`
// 	CreatedAt   time.Time `json:"created_at"`
// 	UpdatedAt   time.Time `json:"updated_at"`
// }
