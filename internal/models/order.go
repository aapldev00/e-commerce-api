package models

import "time"

type Order struct {
	ID         uint        `json:"id"          db:"id"`
	UserID     uint        `json:"user_id"     db:"user_id"`
	TotalPrice float64     `json:"total_price" db:"total_price"`
	Status     string      `json:"status"      db:"status"`
	Items      []OrderItem `json:"items,omitempty" db:"-"`
	CreatedAt  time.Time   `json:"created_at"  db:"created_at"`
}

type OrderItem struct {
	ID        uint    `json:"id"         db:"id"`
	OrderID   uint    `json:"order_id"   db:"order_id"`
	ProductID uint    `json:"product_id" db:"product_id"`
	Quantity  int     `json:"quantity"   db:"quantity"   binding:"required,gt=0"`
	UnitPrice float64 `json:"unit_price" db:"unit_price"`
}
