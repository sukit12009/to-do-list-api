package controllers

import (
	"encoding/json"
	"go-todo-app/config"
	"go-todo-app/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetTasks godoc
// @Summary Get all tasks
// @Description Retrieve all tasks
// @Tags tasks
// @Produce json
// @Success 200 {array} models.Task
// @Router /tasks [get]
func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var tasks []models.Task
	if err := config.DB.Where("deleted_at IS NULL").Find(&tasks).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

// GetOne godoc
// @Summary Get a task by ID
// @Description Retrieve a task from the database by ID
// @Tags tasks
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} models.Task
// @Failure 400 {object} map[string]string "Invalid task ID"
// @Failure 404 {object} map[string]string "Task not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /tasks/{id} [get]
func GetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, `{"error": "Invalid task ID"}`, http.StatusBadRequest)
		return
	}

	var task models.Task
	if err := config.DB.First(&task, id).Error; err != nil {
		http.Error(w, `{"error": "Task not found"}`, http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(task)
}

// CreateTask godoc
// @Summary Create a new task
// @Description Create a new task in the database
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body models.Task true "Task to create"
// @Success 201 {object} models.Task
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /tasks [post]
func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
		return
	}

	if err := config.DB.Create(&task).Error; err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated) // ตั้งสถานะเป็น 201 Created
	json.NewEncoder(w).Encode(task)
}

// UpdateTask godoc
// @Summary Update an existing task
// @Description Update a task in the database by ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param task body models.Task true "Task to update"
// @Success 200 {object} models.Task
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 404 {object} map[string]string "Task not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /tasks/{id} [put]
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, `{"error": "Invalid task ID"}`, http.StatusBadRequest)
		return
	}

	var task models.Task
	if err := config.DB.First(&task, id).Error; err != nil {
		http.Error(w, `{"error": "Task not found"}`, http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
		return
	}

	if err := config.DB.Save(&task).Error; err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(task)
}

// DeleteTask godoc
// @Summary Delete a task
// @Description Delete a task from the database by ID
// @Tags tasks
// @Param id path int true "Task ID"
// @Success 200 {string} string "Task deleted"
// @Failure 400 {object} map[string]string "Invalid task ID"
// @Failure 404 {object} map[string]string "Task not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /tasks/{id} [delete]
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, `{"error": "Invalid task ID"}`, http.StatusBadRequest)
		return
	}

	var task models.Task
	if err := config.DB.First(&task, id).Error; err != nil {
		http.Error(w, `{"error": "Task not found"}`, http.StatusNotFound)
		return
	}

	if err := config.DB.Delete(&task, id).Error; err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Task deleted")
}
