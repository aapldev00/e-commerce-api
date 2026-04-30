package models

import "time"

type Product struct {
	ID          uint       `json:"id"`
	Name        string     `json:"name" binding:"required,min=3"`
	Description string     `json:"description"`
	Price       float64    `json:"price" binding:"required,gt=0"`
	Stock       int        `json:"stock" binding:"min=0"`
	CategoryID  uint       `json:"category_id" binding:"required"`
	Category    Category   `json:"category,omitempty"` // Se llena solo en consultas GET
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"-"`
}
