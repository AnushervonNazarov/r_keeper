package repository

import (
	"r_keeper/db"
	"r_keeper/logger"
	"r_keeper/models"
)

func GetAllOrders(userID uint, query string) ([]models.Order, error) {
	var orders []models.Order

	query = "%" + query + "%"

	if err := db.GetDBConn().Model(&models.Order{}).
		Joins("JOIN users ON users.id = orders.user_id").
		Preload("Items").
		Where("orders.user_id = ? AND description iLIKE ?", userID, query).
		Order("orders.id").
		Find(&orders).Error; err != nil {
		logger.Error.Println("[repository.GetAllOrders] error getting all orders:", err.Error())
		return nil, translateError(err)
	}
	return orders, nil
}

func GetOrderByID(id int) (order models.Order, err error) {

	if err = db.GetDBConn().Where("id = ?", id).First(&order).Error; err != nil {
		logger.Error.Println("[repository.GetOrderByID] error getting order by id. Error is:", err.Error())
		return order, translateError(err)
	}
	return order, nil
}

func CreateOrder(order models.Order) (err error) {
	if err := db.GetDBConn().Create(&order).Error; err != nil {
		logger.Error.Println("[repository.CreateOrder] error creating order. Error is:", err.Error())
		return translateError(err)
	}
	return nil
}

func EditOrderByID(order *models.Order) (*models.Order, error) {
	if err := db.GetDBConn().Save(&order).Error; err != nil {
		logger.Error.Println("[repository.EditOrderByID] error editing order. Error is:", err.Error())
		return nil, translateError(err)
	}
	return order, nil
}

func DeleteOrderByID(order *models.Order) error {
	if err := db.GetDBConn().Delete(order).Error; err != nil {
		logger.Error.Println("[repository.DeleteOrderByID] error deleating order. Error is:", err.Error())
		return translateError(err)
	}
	return nil
}
