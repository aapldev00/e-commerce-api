package models

import "time"

type Product struct {
	ID          uint       `json:"id"          db:"id"`
	Name        string     `json:"name"        db:"name"        binding:"required,min=3"`
	Description string     `json:"description" db:"description"`
	Price       float64    `json:"price"       db:"price"       binding:"required,gt=0"`
	Stock       int        `json:"stock"       db:"stock"       binding:"min=0"`
	CategoryID  uint       `json:"category_id" db:"category_id" binding:"required"`
	Category    *Category  `json:"category,omitempty" db:"-"`
	CreatedAt   time.Time  `json:"created_at"  db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"  db:"updated_at"`
	DeletedAt   *time.Time `json:"-"           db:"deleted_at"`
}
