package models

import "time"

type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Total     float64   `json:"total"`
}

type OrderItem struct {
	ID        int     `json:"id"`
	OrderID   int     `json:"order_id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
