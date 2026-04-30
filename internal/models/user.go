package models

import "time"

type User struct {
	ID        uint      `json:"id"         db:"id"`
	Email     string    `json:"email"      db:"email"      binding:"required,email"`
	Password  string    `json:"password,omitempty" db:"password_hash" binding:"required,min=8"`
	FirstName string    `json:"first_name" db:"first_name" binding:"required"`
	LastName  string    `json:"last_name"  db:"last_name"  binding:"required"`
	Role      string    `json:"role"       db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
