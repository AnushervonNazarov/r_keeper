package models

import "time"

type Report struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	UserID     uint    `json:"user_id" gorm:"not null"`
	User       User    `json:"user" gorm:"foreignKey:UserID;references:ID"`
	TotalSales float64 `json:"total_sales" gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
