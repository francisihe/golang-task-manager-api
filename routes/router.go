package routes

import (
	"net/http"

	"github.com/francisihe/golang-task-manager-api/handlers"
)

// SetupRouter sets up the routes for the API
func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Define the routes
	mux.HandleFunc("/api/tasks", handlers.TaskHandler) // Handle tasks CRUD

	return mux
}
