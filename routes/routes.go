package routes

import (
	"github.com/gorilla/mux"
	"go-todo-app/controllers"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE")

	return router
}
