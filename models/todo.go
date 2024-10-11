package models

import "time"

// Todo represents a Todo item
type Todo struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `gorm:"size:255;not null" validate:"required" json:"title"`
	Description string     `gorm:"type:text;not null" validate:"required" json:"description"`
	CompletedAt *time.Time `gorm:"timestamp" json:"completed_at"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

type UpdateTodoRequest struct {
	Title       string     `json:"title" validate:"required"`
	Description string     `json:"description" validate:"required"`
	CompletedAt *time.Time `json:"completed_at"`
}
