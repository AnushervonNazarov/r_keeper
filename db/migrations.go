package db

import (
	"r_keeper/models"
)

func Migrate() error {
	err := dbConn.AutoMigrate(
		models.User{},
		models.Order{},
		// models.OrderItem{},
		models.Product{})
	if err != nil {
		return err
	}
	return nil
}
