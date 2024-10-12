package handlers

import (
	"fmt"
	"net/http"
)

// TaskHandler handles the requests related to tasks
func TaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTasks(w, r)
	case http.MethodPost:
		createTask(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// getTasks handles GET requests for fetching tasks
func getTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "List of tasks")
}

// createTask handles POST requests for creating a new task
func createTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "New task created")
}
