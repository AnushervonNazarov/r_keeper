package models

import "time"

type Order struct {
	// ID        int       `json:"id"`
	// UserID    int       `json:"user_id"`
	// CreatedAt time.Time `json:"created_at"`
	// Total     float64   `json:"total"`
	ID        uint    `json:"id" gorm:"primaryKey"`
	TableID   uint    `json:"table_id" gorm:"not null"`
	Items     []Menu  `json:"items" gorm:"many2many:order_items;"`
	Total     float64 `json:"total" gorm:"not null"`
	CreatedAt time.Time
}

// type OrderItem struct {
// 	ID        int     `json:"id"`
// 	OrderID   int     `json:"order_id"`
// 	ProductID int     `json:"product_id"`
// 	Quantity  int     `json:"quantity"`
// 	Price     float64 `json:"price"`
// }
