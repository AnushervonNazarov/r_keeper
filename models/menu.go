package models

import "time"

// type Menu struct {
// 	ID       uint    `json:"id" gorm:"primaryKey"`
// 	Name     string  `json:"name" gorm:"not null"`
// 	Price    float64 `json:"price" gorm:"not null"`
// 	Category string  `json:"category" gorm:"not null"`
// }

type Menu struct {
	ID        uint        `json:"id" gorm:"primaryKey"`
	Name      string      `json:"name" gorm:"not null"`
	Price     float64     `json:"price" gorm:"not null"`
	Category  string      `json:"category" gorm:"not null"`
	Items     []OrderItem `json:"items" gorm:"foreignKey:MenuItemID;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SwagMenu struct {
	ID       uint    `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name" gorm:"not null"`
	Price    float64 `json:"price" gorm:"not null"`
	Category string  `json:"category" gorm:"not null"`
}
