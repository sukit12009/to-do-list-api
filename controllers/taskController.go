package controllers

import (
	"encoding/json"
	"go-todo-app/config"
	"go-todo-app/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	config.DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	config.DB.Create(&task)
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var task models.Task
	config.DB.First(&task, id)
	json.NewDecoder(r.Body).Decode(&task)
	config.DB.Save(&task)
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var task models.Task
	config.DB.Delete(&task, id)
	json.NewEncoder(w).Encode("Task deleted")
}
