package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/francisihe/golang-task-manager-api/models"
)

var tasks []models.Task

// Helper function to generate a random task ID
func generateID() string {
	return fmt.Sprintf("%d", rand.Intn(100000))
}

// TaskHandler handles the requests related to tasks
func TaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet: // Equivalent to "GET"
		// You could use "GET" as well, but http.MethodGet is preferred for consistency and to avoid typos
		GetTasks(w, r)

	case http.MethodPost: // Equivalent to "POST"
		// You could use "POST", but http.MethodPost ensures consistency with Go's http package constants
		CreateTask(w, r)

	case http.MethodPut: // Equivalent to "PUT"
		// You could use "PUT", but http.MethodPut is more robust and error-proof
		UpdateTask(w, r)

	case http.MethodDelete: // Equivalent to "DELETE"
		// You could use "DELETE", but http.MethodDelete ensures clarity and prevents potential string errors
		DeleteTask(w, r)

	default:
		// Handle unsupported methods
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// In essence, the above TaskHandler function could still be written as this below:

// func TaskHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case "GET":
// 		GetTasks(w, r)
// 	case "POST":
// 		CreateTask(w, r)
// 	case "PUT":
// 		UpdateTask(w, r)
// 	case "DELETE":
// 		DeleteTask(w, r)
// 	default:
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 	}
// }

// // TaskHandler handles the requests related to tasks -- A simulated version
// // getTasks handles GET requests for fetching tasks
// func getTasks(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "List of tasks")
// }

// // createTask handles POST requests for creating a new task
// func createTask(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "New task created")
// }

// CreateTask handles the creation of a new task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask models.Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newTask.ID = generateID() // Generate random ID
	newTask.CreatedAt = time.Now()
	newTask.UpdatedAt = time.Now()
	tasks = append(tasks, newTask)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

// GetTasks retrieves all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// UpdateTask handles updating an existing task
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var updatedTask models.Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, task := range tasks {
		if task.ID == updatedTask.ID {
			updatedTask.UpdatedAt = time.Now()
			tasks[i] = updatedTask
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updatedTask)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

// DeleteTask handles the deletion of a task
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	var taskToDelete models.Task
	if err := json.NewDecoder(r.Body).Decode(&taskToDelete); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, task := range tasks {
		if task.ID == taskToDelete.ID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(taskToDelete)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}
