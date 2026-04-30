package models

import "time"

type Category struct {
	ID          uint      `json:"id"          db:"id"`
	Name        string    `json:"name"        db:"name"        binding:"required"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at"  db:"created_at"`
}
