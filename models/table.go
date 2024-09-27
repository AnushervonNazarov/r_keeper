package models

import "time"

// type Table struct {
// 	ID          int       `json:"id" gorm:"primaryKey"`
// 	TableNumber int       `json:"table_number" gorm:"not null"`
// 	Reserved    bool      `json:"reserved" gorm:"default:false"`
// 	Capacity    int       `json:"capacity" gorm:"not null"`
// 	Orders      []Order   `json:"orders" gorm:"foreignKey:TableID;constraint:OnDelete:CASCADE;"` // Связь один ко многим с Order
// 	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
// 	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
// }

type Table struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	TableNumber int       `json:"table_number" gorm:"not null"`
	Reserved    bool      `json:"reserved" gorm:"default:false"`
	Capacity    int       `json:"capacity" gorm:"not null"`
	Orders      []Order   `json:"orders" gorm:"foreignKey:TableID;constraint:OnDelete:CASCADE;"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type TableResponse struct {
	ID          int             `json:"id"`
	TableNumber int             `json:"table_number"`
	Reserved    bool            `json:"reserved"`
	Capacity    int             `json:"capacity"`
	Orders      []OrderResponse `json:"orders"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

type SwagTable struct {
	ID          int  `json:"id" gorm:"primaryKey"`
	TableNumber int  `json:"table_number" gorm:"not null"`
	Reserved    bool `json:"reserved" gorm:"default:false"`
	Capacity    int  `json:"capacity" gorm:"not null"`
}
