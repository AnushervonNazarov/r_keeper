package models

import "time"

type Role struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"not null"`
	Users     []User `json:"users" gorm:"foreignKey:RoleID;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
