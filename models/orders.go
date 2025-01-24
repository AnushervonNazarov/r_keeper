package models

import "time"

type Order struct {
	ID          int         `json:"id" gorm:"primaryKey"`
	TableID     int         `json:"table_id" gorm:"not null"`
	UserID      string      `json:"user_id" gorm:"not null"`
	TotalAmount float64     `json:"total_amount" gorm:"not null"`
	Items       []OrderItem `json:"items" gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;"`
	CreatedAt   time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
}

type OrderItem struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	OrderID    int       `json:"order_id" gorm:"not null"`
	MenuItemID int       `json:"menu_item_id" gorm:"not null"`
	MenuItem   Menu      `json:"menu_item" gorm:"foreignKey:MenuItemID;constraint:OnDelete:CASCADE;"`
	Quantity   int       `json:"quantity" gorm:"not null"`
	Price      float64   `json:"price" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type OrderResponse struct {
	ID          int            `json:"id"`
	TableID     int            `json:"table_id"`
	UserID      string         `json:"user_id"`
	TotalAmount float64        `json:"total_amount"`
	Items       []OrderItemDTO `json:"items"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type OrderItemDTO struct {
	ID         int       `json:"id"`
	MenuItemID int       `json:"menu_item_id"`
	Quantity   int       `json:"quantity"`
	Price      float64   `json:"price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
type SwagOrder struct {
	TableID     int         `json:"table_id" gorm:"not null"`
	UserID      int         `json:"user_id" gorm:"not null"`
	Items       []OrderItem `json:"items" gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;"`
	TotalAmount float64     `json:"total_amount" gorm:"not null"`
}
