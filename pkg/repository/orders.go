package repository

import (
	"errors"
	"r_keeper/db"
	"r_keeper/logger"
	"r_keeper/models"
)

func GetAllOrders() (orders []models.Order, err error) {
	if err = db.GetDBConn().Preload("Items").Find(&orders).Error; err != nil {
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

func GetAllChecks() (checks []models.Check, err error) {
	if err = db.GetDBConn().Find(&checks).Error; err != nil {
		logger.Error.Println("[repository.GetAllChecks] error getting all checks:", err.Error())
		return nil, translateError(err)
	}
	return checks, nil
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

// Сохранение чека в базе данных
func SaveCheck(check models.Check) error {
	// Здесь будет логика сохранения чека в базу данных
	if err := db.GetDBConn().Save(&check).Error; err != nil {
		logger.Error.Println("[repository.SaveCheck] error saving check. Error is:", err.Error())
		return translateError(err)
	}
	// Например, SQL-запрос или сохранение в файл

	// Пример обработки ошибки
	if check.OrderID == 0 {
		return errors.New("invalid order ID")
	}

	// Заглушка для успешного сохранения
	return nil
}

// Получение чека по ID
func GetCheckByID(id int) (check models.Check, err error) {
	// Логика получения чека из базы данных
	if err = db.GetDBConn().Where("id = ?", id).Preload("Items").First(&check).Error; err != nil {
		logger.Error.Println("[repository.GetCheckByID] error getting check by id. Error is:", err.Error())
		return check, translateError(err)
	}
	// Пример заглушки
	return check, nil
}

func GetOrderByIDForReceipt(orderID int) (models.Order, error) {
	var order models.Order
	err := db.GetDBConn().Preload("Items.MenuItem").First(&order, orderID).Error
	if err != nil {
		logger.Error.Println("[repository.GetOrderByID] error getting order by id. Error is:", err.Error())
		return order, translateError(err)
	}
	return order, nil
}
