package models

type Dishes struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description" gorm:"not null"`
	Price       float64 `json:"price" gorm:"not null"`
	Stock       int     `json:"stock" gorm:"not null"`
}
