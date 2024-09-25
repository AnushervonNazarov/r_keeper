package models

import "time"

type Table struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	TableNumber int       `json:"table_number" gorm:"not null"`
	Reserved    bool      `json:"reserved" gorm:"default:false"`
	Capacity    int       `json:"capacity" gorm:"not null"`
	Orders      []Order   `json:"orders" gorm:"foreignKey:TableID;constraint:OnDelete:CASCADE;"` // Связь один ко многим с Order
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type SwagTable struct {
	ID          uint `json:"id" gorm:"primaryKey"`
	TableNumber int  `json:"table_number" gorm:"not null"`
	Reserved    bool `json:"reserved" gorm:"default:false"`
	Capacity    int  `json:"capacity" gorm:"not null"`
}

// type ReserveTable struct {
// 	ID       uint `json:"id" gorm:"primaryKey"`
// 	Reserved bool `json:"reserved"`
// }
