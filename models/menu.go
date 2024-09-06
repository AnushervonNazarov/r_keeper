package models

type Menu struct {
	ID       uint    `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name" gorm:"not null"`
	Price    float64 `json:"price" gorm:"not null"`
	Category string  `json:"category" gorm:"not null"`
}
