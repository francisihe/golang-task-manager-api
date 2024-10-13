package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/francisihe/golang-task-manager-api/models"
	"github.com/francisihe/golang-task-manager-api/routes"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Print Hello World message
	helloWorld()

	// ========================
	// Initialize the database connection
	dsn := "host=localhost user=postgres password=yourpassword dbname=golang_task_manager port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Perform auto-migration to create/update tables based on the models
	db.AutoMigrate(&models.Task{})

	// Set the database instance globally -- Instantiated in the models package
	models.DB = db
	// ========================

	// Initialize the router
	router := routes.SetupRouter()

	// Start the HTTP server
	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

func helloWorld() {
	fmt.Println("Hello World\nFrancis is trying out Go!")
}
