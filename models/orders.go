package models

import "time"

type Order struct {
	ID          uint        `json:"id" gorm:"primaryKey"`
	TableID     uint        `json:"table_id" gorm:"not null"`                                     // Внешний ключ на таблицу Table
	Table       Table       `json:"table" gorm:"foreignKey:TableID;constraint:OnDelete:CASCADE;"` // Связь с Table
	UserID      uint        `json:"user_id" gorm:"not null"`                                      // Внешний ключ на таблицу User
	User        User        `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`   // Связь с User
	TotalAmount float64     `json:"total_amount" gorm:"not null"`
	Items       []OrderItem `json:"items" gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;"` // Связь один ко многим с OrderItem
	CreatedAt   time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
}

type OrderItem struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	OrderID    uint      `json:"order_id" gorm:"not null"`                                            // Внешний ключ на таблицу Order
	Order      Order     `json:"order" gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;"`        // Связь с Order
	MenuItemID uint      `json:"menu_item_id" gorm:"not null"`                                        // Внешний ключ на таблицу Menu
	MenuItem   Menu      `json:"menu_item" gorm:"foreignKey:MenuItemID;constraint:OnDelete:CASCADE;"` // Связь с Menu
	Quantity   int       `json:"quantity" gorm:"not null"`
	Price      float64   `json:"price" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// type Inventory struct {
// 	ID        uint    `json:"id" gorm:"primaryKey"`
// 	ItemName  string  `json:"item_name" gorm:"not null"`
// 	Quantity  float64 `json:"quantity" gorm:"not null"`
// 	Unit      string  `json:"unit" gorm:"not null"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

type SwagOrder struct {
	Table    SwagTable `json:"table" gorm:"foreignKey:TableID;references:ID;constraint:OnDelete:CASCADE;"`
	User     SwagUser  `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	Quantity int       `json:"qunatity" gorm:"not null"`
	Price    float64   `json:"price" gorm:"not null"`
}
