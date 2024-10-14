package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// middleware "github.com/francisihe/golang-task-manager-api/middlewares" -- importing a packaga as 'another name' syntax
	"github.com/francisihe/golang-task-manager-api/models"
	"github.com/francisihe/golang-task-manager-api/routes"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Function to load the environment variables
func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	// Print Hello World message
	helloWorld()

	// DATABASE VARIABLES
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// ========================
	// Initialize the database connection
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	// dsn := "host=localhost user=postgres password=yourpassword dbname=golang_task_manager port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Initialize your application with the database connection
	log.Println("Database connected successfully")

	// Perform auto-migration to create/update tables based on the models
	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatalf("Error during migration: %v", err)
	}

	// Set the database instance globally -- Instantiated in the models package
	models.DB = db
	// ========================

	// Initialize the router
	router := routes.SetupRouter()

	// ==== ROUTING USING GORILLA MUX / MIDDLEWARE APPLICATION ===

	// If i were using the gorilla mux router, it would look like this:
	// Use the mux router from routes.SetupRouter
	// http.Handle("/", routes.SetupRouter())

	// And apply the middleware as such:
	// // Apply rate limiting middleware
	// http.Handle("/", middleware.RateLimiting(mux))

	// Apply rate limiting middleware to the router using gorilla/mux
	// // Apply the middleware globally using mux.Use
	// // Apply CORS handling
	// router.Use(middleware.CORS)

	// // Apply rate limiting
	// router.Use(middleware.RateLimiting)

	// ===========================

	// Start the HTTP server
	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

func helloWorld() {
	fmt.Println("Hello World\nFrancis is trying out Go!")
}
