package repository

import (
	"r_keeper/db"
	"r_keeper/logger"
	"r_keeper/models"
)

func GetAllOrders() (orders []models.Order, err error) {
	if err = db.GetDBConn().Preload("Items").Find(&orders).Error; err != nil {
		logger.Error.Printf("[repository.GetAllOrders] error getting all orders: %s\n", err.Error())
		return nil, err
	}
	return orders, nil
}

func GetOrderByID(id int) (order models.Order, err error) {

	if err = db.GetDBConn().Where("id = ?", id).First(&order).Error; err != nil {
		logger.Error.Printf("[repository.GetOrderByID] error getting order by id: %v\n", err)
		return order, err
	}
	return order, nil
}

func CreateOrder(order models.Order) (err error) {
	if err := db.GetDBConn().Create(&order).Error; err != nil {
		logger.Error.Printf("[repository.CreateOrder] error creating order: %v\n", err)
		return err
	}
	return nil
}

func EditOrderByID(order *models.Order) (*models.Order, error) {
	if err := db.GetDBConn().Save(&order).Error; err != nil {
		logger.Error.Printf("[repository.EditOrderByID] error editing order: %v\n", err)
		return nil, err
	}
	return order, nil
}

func DeleteOrderByID(order *models.Order) error {
	if err := db.GetDBConn().Delete(order).Error; err != nil {
		logger.Error.Printf("[repository.DeleteOrderByID] error deleating order: %v\n", err)
		return err
	}
	return nil
}
