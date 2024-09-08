package models

import "time"

type Table struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	TableNumber int     `json:"table_number" gorm:"not null"`
	Capacity    int     `json:"capacity" gorm:"not null"`
	Orders      []Order `json:"orders" gorm:"foreignKey:TableID;references:ID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
