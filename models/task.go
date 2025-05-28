package models

import (
	"go-todo-app/config"

	"time"
)

type Task struct {
	ID        uint       `json:"id" example:"1"`
	Title     string     `json:"title" example:"Buy groceries"`
	Completed bool       `json:"completed" example:"false"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func AutoMigrate() {
	config.DB.AutoMigrate(&Task{})
}
