package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/francisihe/golang-task-manager-api/models"
)

// Helper function to generate a random task ID
func generateID() string {
	return fmt.Sprintf("%d", rand.Intn(100000))
}

// TaskHandler handles the requests related to tasks
func TaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetTasks(w, r)

	case http.MethodPost:
		CreateTask(w, r)

	case http.MethodPatch:
		UpdateTask(w, r)

	case http.MethodDelete:
		DeleteTask(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// -------------------------------------------------------------------------
// New GORM-based handlers interacting with PostgreSQL

// CreateTask handles the creation of a new task with GORM
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask models.Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate random ID, set timestamps
	newTask.ID = generateID()
	newTask.CreatedAt = time.Now()
	newTask.UpdatedAt = time.Now()

	// Save task to the database using GORM
	if err := models.DB.Create(&newTask).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

// GetTasks retrieves all tasks from the database using GORM
func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

	// Fetch tasks from the database
	if err := models.DB.Find(&tasks).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// UpdateTask handles updating an existing task with GORM
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var updatedTask models.Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Find task by ID and update
	var task models.Task
	if err := models.DB.First(&task, "id = ?", updatedTask.ID).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	// Update fields and save to the database
	task.Title = updatedTask.Title
	task.Description = updatedTask.Description
	task.Status = updatedTask.Status
	task.UpdatedAt = time.Now()

	if err := models.DB.Save(&task).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

// DeleteTask handles the deletion of a task with GORM
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	var taskToDelete models.Task
	if err := json.NewDecoder(r.Body).Decode(&taskToDelete); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Find task by ID and delete from the database
	if err := models.DB.Where("id = ?", taskToDelete.ID).Delete(&models.Task{}).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(taskToDelete)
}

// -------------------------------------------------------------------------
// Old in-memory task handling logic (commented out for reference)

// CreateTask handles the creation of a new task
// func CreateTask(w http.ResponseWriter, r *http.Request) {
// 	var newTask models.Task
// 	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
//
// 	newTask.ID = generateID() // Generate random ID
// 	newTask.CreatedAt = time.Now()
// 	newTask.UpdatedAt = time.Now()
// 	tasks = append(tasks, newTask)
//
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(newTask)
// }

// GetTasks retrieves all tasks
// func GetTasks(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(tasks)
// }

// UpdateTask handles updating an existing task
// func UpdateTask(w http.ResponseWriter, r *http.Request) {
// 	var updatedTask models.Task
// 	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
//
// 	for i, task := range tasks {
// 		if task.ID == updatedTask.ID {
// 			updatedTask.UpdatedAt = time.Now()
// 			tasks[i] = updatedTask
// 			w.WriteHeader(http.StatusOK)
// 			json.NewEncoder(w).Encode(updatedTask)
// 			return
// 		}
// 	}
// 	http.Error(w, "Task not found", http.StatusNotFound)
// }

// DeleteTask handles the deletion of a task
// func DeleteTask(w http.ResponseWriter, r *http.Request) {
// 	var taskToDelete models.Task
// 	if err := json.NewDecoder(r.Body).Decode(&taskToDelete); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
//
// 	for i, task := range tasks {
// 		if task.ID == taskToDelete.ID {
// 			tasks = append(tasks[:i], tasks[i+1:]...)
// 			w.WriteHeader(http.StatusOK)
// 			json.NewEncoder(w).Encode(taskToDelete)
// 			return
// 		}
// 	}
// 	http.Error(w, "Task not found", http.StatusNotFound)
// }
