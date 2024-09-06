package models

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"user_name" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Role     string `json:"role" gorm:"not null"` // 'admin', 'waiter'
}
