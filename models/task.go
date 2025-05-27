package models

import (
	"go-todo-app/config"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func AutoMigrate() {
	config.DB.AutoMigrate(&Task{})
}
