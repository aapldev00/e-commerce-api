package models

import "time"

type User struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email" binding:"required,email"`
	Password  string    `json:"password,omitempty" binding:"required,min=8"`
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
