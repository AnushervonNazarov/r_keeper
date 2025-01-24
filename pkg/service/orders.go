package service

import (
	"errors"
	"fmt"
	"r_keeper/errs"
	"r_keeper/models"
	"r_keeper/pkg/repository"
	"strconv"
	"strings"
	"time"
)

func GetAllOrders() ([]models.OrderResponse, error) {
	var orders []models.Order
	err := repository.GetAllOrders(&orders)
	if err != nil {
		return nil, err
	}

	var orderResponses []models.OrderResponse
	for _, order := range orders {
		var orderResponse models.OrderResponse
		orderResponse.ID = order.ID
		orderResponse.TableID = order.TableID
		orderResponse.UserID = order.UserID
		orderResponse.TotalAmount = order.TotalAmount
		orderResponse.CreatedAt = order.CreatedAt
		orderResponse.UpdatedAt = order.UpdatedAt

		for _, item := range order.Items {
			orderResponse.Items = append(orderResponse.Items, models.OrderItemDTO{
				ID:         item.ID,
				MenuItemID: item.MenuItemID,
				Quantity:   item.Quantity,
				Price:      item.Price,
				CreatedAt:  item.CreatedAt,
				UpdatedAt:  item.UpdatedAt,
			})
		}
		orderResponses = append(orderResponses, orderResponse)
	}

	return orderResponses, nil
}

func GetOrdersForUser(userID string) ([]models.Order, error) {
	return repository.GetOrdersByUserID(userID)
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

	orderInput.ID = id

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

func CreateCheck(orderID int, tableNumber int, items []models.CheckItem) (check models.Check, err error) {
	check = models.Check{
		OrderID:     orderID,
		TableNumber: tableNumber,
		Items:       items,
	}

	check.CalculateTotal()

	err = repository.SaveCheck(check)
	if err != nil {
		return check, errors.New("failed to save check")
	}

	return check, nil
}

func GenerateReceipt(orderID int, commissionRate float64) (string, error) {
	order, err := repository.GetOrderByIDForReceipt(orderID)
	if err != nil {
		return "", fmt.Errorf("failed to fetch order: %w", err)
	}

	table, err := repository.GetTableByIDForReceipt(order.TableID)
	if err != nil {
		return "", fmt.Errorf("failed to fetch table: %w", err)
	}

	var receiptBuilder strings.Builder
	receiptBuilder.WriteString("========== Receipt ==========\n")
	receiptBuilder.WriteString("Order number: " + strconv.Itoa(order.ID) + "\n")
	receiptBuilder.WriteString("Table: " + strconv.Itoa(table.TableNumber) + "\n")
	receiptBuilder.WriteString("Data: " + time.Now().Format("02-01-2006 15:04") + "\n")
	receiptBuilder.WriteString("=========================\n")

	receiptBuilder.WriteString("Name     Quantity     Price     Sum\n")
	receiptBuilder.WriteString("-------------------------------------------\n")
	var total float64
	for _, item := range order.Items {
		itemTotal := float64(item.Quantity) * item.Price
		total += itemTotal
		receiptBuilder.WriteString(fmt.Sprintf("%-16s %-10d %-8.2f %-8.2f\n", item.MenuItem.Name, item.Quantity, item.Price, itemTotal))
	}

	commission := total * commissionRate
	netTotal := total + commission

	receiptBuilder.WriteString("-------------------------------------------\n")
	receiptBuilder.WriteString(fmt.Sprintf("Commission (%.0f%%): %25.2f\n", commissionRate*100, commission))
	receiptBuilder.WriteString(fmt.Sprintf("Summary :             %23.2f\n", netTotal))
	receiptBuilder.WriteString("=========================\n")

	return receiptBuilder.String(), nil
}
