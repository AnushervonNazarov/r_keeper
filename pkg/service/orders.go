package service

import (
	"errors"
	"fmt"
	"r_keeper/errs"
	"r_keeper/models"
	"r_keeper/pkg/repository"
)

func GetAllOrders() (orders []models.Order, err error) {
	if orders, err = repository.GetAllOrders(); err != nil {
		return nil, err
	}
	return orders, nil
}

func GetOrderByID(id int) (order models.Order, err error) {
	if order, err = repository.GetOrderByID(id); err != nil {
		return order, err
	}
	return order, nil
}

func GetAllChecks() (checks []models.Check, err error) {
	if checks, err = repository.GetAllChecks(); err != nil {
		return nil, err
	}
	return checks, nil
}

func GetCheckByID(id int) (check models.Check, err error) {
	if check, err = repository.GetCheckByID(id); err != nil {
		return check, err
	}
	return check, nil
}

func CreateOrder(order models.Order) error {
	_, err := repository.GetOrderByID(int(order.ID))
	if err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		return err
	}

	order.Table.Reserved = true

	if err = repository.CreateOrder(order); err != nil {
		return err
	}
	return nil
}

func EditOrderByID(id int, orderInput models.Order) (*models.Order, error) {
	_, err := repository.GetOrderByID(id)
	if err != nil {
		return nil, fmt.Errorf("order not found: %v", err)
	}

	orderInput.ID = uint(id)

	updatedOrder, err := repository.EditOrderByID(&orderInput)
	if err != nil {
		return nil, fmt.Errorf("could not update order: %v", err)
	}

	return updatedOrder, nil
}

func DeleteOrderByID(id int) error {
	order, err := repository.GetOrderByID(id)
	if err != nil {
		return fmt.Errorf("order not found: %v", err)
	}

	if err := repository.DeleteOrderByID(&order); err != nil {
		return fmt.Errorf("could not delete order: %v", err)
	}

	return nil
}

// Создание чека
func CreateCheck(orderID int, tableNumber int, items []models.CheckItem) (check models.Check, err error) {
	check = models.Check{
		OrderID:     orderID,
		TableNumber: tableNumber,
		Items:       items,
	}

	check.CalculateTotal()

	// Сохраняем чек через репозиторий
	err = repository.SaveCheck(check)
	if err != nil {
		return check, errors.New("failed to save check")
	}

	return check, nil
}
