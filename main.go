package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go-todo-app/config"
	"go-todo-app/middleware"
	"go-todo-app/models"
	"go-todo-app/routes"

	_ "go-todo-app/docs"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	config.ConnectDatabase()
	models.AutoMigrate()

	r := middleware.CORS(routes.RegisterRoutes())

	fmt.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
