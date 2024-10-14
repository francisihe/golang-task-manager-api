package routes

import (
	"net/http"

	"github.com/francisihe/golang-task-manager-api/handlers"
	middleware "github.com/francisihe/golang-task-manager-api/middlewares"
)

// SetupRouter sets up the routes for the API
func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Define the routes
	mux.HandleFunc("/api/tasks", handlers.TaskHandler) // Handle tasks CRUD

	mux.HandleFunc("/api/login", handlers.Login) // Login for generating JWT

	// Apply JWT authentication middleware to task-related routes
	mux.Handle("/api/tasks/", middleware.AuthMiddleware(http.StripPrefix("/api", mux)))

	// Alternatively, I could define the routes individually like so:
	// mux.HandleFunc("/api/tasks/create", handlers.CreateTask)
	// mux.HandleFunc("/api/tasks/get", handlers.GetTasks)
	// mux.HandleFunc("/api/tasks/update", handlers.UpdateTask)
	// mux.HandleFunc("/api/tasks/delete", handlers.DeleteTask)

	// -- Alternatively, and ideally, I prefer to specify the methods in the routes.
	// -- This way, I need to use the gorilla/mux package to handle the routes.
	// -- This is because the standard http.ServeMux does not support specifying methods in the routes.
	// -- The gorilla/mux package is more powerful and flexible than the standard http.ServeMux.
	// -- I need to also add the router to the main.go file.

	return mux
}

// If i were using the gorilla mux router, it would look like this:

// package routes

// import (
// 	"net/http"
// 	"github.com/gorilla/mux"
// 	"github.com/francisihe/golang-task-manager-api/handlers"
// )

// // SetupRouter sets up the routes for the API
// func SetupRouter() *mux.Router {
// 	r := mux.NewRouter()

// 	// Define routes with HTTP methods
// 	r.HandleFunc("/api/tasks", handlers.GetTasks).Methods("GET") // Get tasks
// 	r.HandleFunc("/api/tasks", handlers.CreateTask).Methods("POST") // Create task
// 	r.HandleFunc("/api/tasks/{id}", handlers.UpdateTask).Methods("PUT") // Update task
// 	r.HandleFunc("/api/tasks/{id}", handlers.DeleteTask).Methods("DELETE") // Delete task

// 	return r
// }
