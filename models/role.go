package models

import "time"

type Role struct {
	ID        uint      `json:"-" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Users     []User    `json:"-" gorm:"foreignKey:RoleID;references:ID"`
	CreatedAt time.Time `json:"created_at" gorm:"autocreatedtime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoupdatedtime"`
}
