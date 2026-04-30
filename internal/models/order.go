package models

import "time"

type Order struct {
	ID         uint        `json:"id"`
	UserID     uint        `json:"user_id"`
	TotalPrice float64     `json:"total_price"`
	Status     string      `json:"status"`
	Items      []OrderItem `json:"items,omitempty"`
	CreatedAt  time.Time   `json:"created_at"`
}

type OrderItem struct {
	ID        uint    `json:"id"`
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity" binding:"required,gt=0"`
	UnitPrice float64 `json:"unit_price"`
}
