package models

import (
	"time"

	"gorm.io/gorm"
)

type Check struct {
	ID          int         `json:"id" gorm:"primaryKey"`
	OrderID     int         `json:"order_id"`
	TableNumber int         `json:"table_number"`
	DateTime    time.Time   `json:"datetime"`
	Items       []CheckItem `json:"items" gorm:"foreignKey:CheckID"`
	TotalAmount float64     `json:"total_amount"`
	Tax         float64     `json:"tax"`
}

type CheckItem struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	CheckID   uint           `json:"check_id"`
	Name      string         `json:"name"`
	Quantity  int            `json:"quantity"`
	Price     float64        `json:"price"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (c *Check) CalculateTotal() {
	var total float64
	for _, item := range c.Items {
		total += item.Price * float64(item.Quantity)
	}
	c.TotalAmount = total
	c.Tax = total * 0.1 // предположим, что налог составляет 10%
}
