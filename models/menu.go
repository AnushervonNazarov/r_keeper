package models

import "time"

// type Menu struct {
// 	ID        uint        `json:"id" gorm:"primaryKey"`
// 	Name      string      `json:"name" gorm:"not null"`
// 	Price     float64     `json:"price" gorm:"not null"`
// 	Category  string      `json:"category" gorm:"not null"`
// 	Items     []OrderItem `json:"items" gorm:"foreignKey:MenuItemID"` // Связь один ко многим с OrderItem
// 	CreatedAt time.Time   `json:"created_at" gorm:"autoCreateTime"`
// 	UpdatedAt time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
// }

type Menu struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Price     float64   `json:"price" gorm:"not null"`
	Category  string    `json:"category" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type MenuResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SwagMenu struct {
	ID       uint    `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name" gorm:"not null"`
	Price    float64 `json:"price" gorm:"not null"`
	Category string  `json:"category" gorm:"not null"`
}
