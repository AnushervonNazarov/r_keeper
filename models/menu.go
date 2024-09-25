package models

import "time"

type Menu struct {
	ID        uint        `json:"id" gorm:"primaryKey"`
	Name      string      `json:"name" gorm:"not null"`
	Price     float64     `json:"price" gorm:"not null"`
	Category  string      `json:"category" gorm:"not null"`
	Items     []OrderItem `json:"items" gorm:"foreignKey:MenuItemID;constraint:OnDelete:CASCADE;"` // Связь один ко многим с OrderItem
	CreatedAt time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
}

type SwagMenu struct {
	ID       uint    `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name" gorm:"not null"`
	Price    float64 `json:"price" gorm:"not null"`
	Category string  `json:"category" gorm:"not null"`
}
