package models

import (
	"time"
)

// type User struct {
// 	ID       int    `json:"id" gorm:"primaryKey"`
// 	Username string `json:"user_name" gorm:"unique;not null"`
// 	Password string `json:"password" gorm:"not null"`
// 	Role     string `json:"role" gorm:"not null"` // 'admin', 'waiter'
// }

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password" gorm:"not null"`
	// Role      Role      `json:"-" gorm:"foreignKey:RoleID;references:ID"`
	// RoleID    uint      `json:"-" gorm:"not null"`
	Role      string    `json:"role" gorm:"not null"`
	IsDeleted bool      `json:"is_deleated" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autocreatedtime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoupdatedtime"`
}

type SwagUser struct {
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password" gorm:"not null"`
}

type SignInInput struct {
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password" gorm:"not null"`
}
