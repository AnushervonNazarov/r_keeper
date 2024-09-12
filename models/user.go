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
	ID        uint   `json:"id" gorm:"primaryKey"`
	Username  string `json:"username" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	RoleID    uint   `json:"role_id" gorm:"not null"`
	Role      Role   `json:"role" gorm:"foreignKey:RoleID;references:ID"`
	IsDeleted bool   `json:"is_deleted" gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SwagUser struct {
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password" gorm:"not null"`
}

type SignInInput struct {
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password" gorm:"not null"`
}
