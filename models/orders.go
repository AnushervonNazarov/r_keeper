package models

import "time"

// type Order struct {
// 	// ID        int       `json:"id"`
// 	// UserID    int       `json:"user_id"`
// 	// CreatedAt time.Time `json:"created_at"`
// 	// Total     float64   `json:"total"`
// 	ID        uint        `json:"id" gorm:"primaryKey"`
// 	TableID   uint        `json:"table_id" gorm:"not null"`
// 	Items     []OrderItem `json:"items" gorm:"many2many:order_items;"`
// 	Total     float64     `json:"total" gorm:"not null"`
// 	CreatedAt time.Time
// }

// type OrderItem struct {
// 	ID        int     `json:"id"`
// 	OrderID   int     `json:"order_id"`
// 	ProductID int     `json:"product_id"`
// 	Quantity  int     `json:"quantity"`
// 	Price     float64 `json:"price"`
// }

type Order struct {
	ID          uint        `json:"id" gorm:"primaryKey"`
	TableID     uint        `json:"table_id" gorm:"not null"`
	Table       Table       `json:"table" gorm:"foreignKey:TableID;references:ID;constraint:OnDelete:CASCADE;"`
	UserID      uint        `json:"user_id" gorm:"not null"`
	User        User        `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	TotalAmount float64     `json:"total_amount" gorm:"not null"`
	Items       []OrderItem `json:"items" gorm:"foreignKey:OrderID;references:ID;constraint:OnDelete:CASCADE;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type OrderItem struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	OrderID    uint    `json:"order_id" gorm:"not null"`
	Order      Order   `json:"order" gorm:"foreignKey:OrderID;references:ID;constraint:OnDelete:CASCADE;"`
	MenuItemID uint    `json:"menu_item_id" gorm:"not null"`
	MenuItem   Menu    `json:"menu_item" gorm:"foreignKey:MenuItemID;references:ID"`
	Quantity   int     `json:"qunatity" gorm:"not null"`
	Price      float64 `json:"price" gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Inventory struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	ItemName  string  `json:"item_name" gorm:"not null"`
	Quantity  float64 `json:"quantity" gorm:"not null"`
	Unit      string  `json:"unit" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SwagOrder struct {
	Table    uint `json:"table" gorm:"foreignKey:TableID;references:ID;constraint:OnDelete:CASCADE;"`
	User     uint `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	Quantity int  `json:"qunatity" gorm:"not null"`
}
