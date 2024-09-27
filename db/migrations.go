package db

import (
	"r_keeper/models"
)

func Migrate() error {
	err := dbConn.AutoMigrate(
		models.Table{},
		models.User{},
		models.Order{},
		models.OrderItem{})
	if err != nil {
		return err
	}
	return nil
}
