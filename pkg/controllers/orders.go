package controllers

import (
	"net/http"
	"r_keeper/errs"
	"r_keeper/logger"
	"r_keeper/models"
	"r_keeper/pkg/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllOrders
// @Summary Get All Orders
// @Security ApiKeyAuth
// @Tags orders
// @Description get list of orders
// @ID get-all-orders
// @Produce json
// @Param q query string false "fill if you need search"
// @Success 200 {array} models.SwagOrder
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/orders [get]
func GetAllOrders(c *gin.Context) {
	userRole := c.GetString(userRoleCtx)
	if userRole != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	orders, err := service.GetAllOrders()
	if err != nil {
		logger.Error.Printf("[controllers.GetAllOrders] error getting all orders: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})
}

// GetOrderByID
// @Summary Get Order By ID
// @Security ApiKeyAuth
// @Tags orders
// @Description get order by ID
// @ID get-order-by-id
// @Produce json
// @Param id path integer true "id of the order"
// @Success 200 {object} models.SwagOrder
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/orders/{id} [get]
func GetOrderByID(c *gin.Context) {
	userRole := c.GetString(userRoleCtx)
	if userRole != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetOrderByID] error getting order %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})

		return
	}

	order, err := service.GetOrderByID(id)
	if err != nil {
		logger.Error.Printf("[controllers.GetOrderByID] error getting order %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, order)
}

// CreateOrder
// @Summary Create Order
// @Security ApiKeyAuth
// @Tags orders
// @Description create new order
// @ID create-order
// @Accept json
// @Produce json
// @Param input body models.SwagOrder true "new order info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/orders [post]
func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		logger.Error.Printf("[controllers.CreateOrder] error creating order %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	err := service.CreateOrder(order)
	if err != nil {
		logger.Error.Printf("[controllers.CreateOrder] error creating order %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "order created successfully",
	})
}

// EditOrderByID
// @Summary Edit Order
// @Security ApiKeyAuth
// @Tags orders
// @Description edit existed order
// @ID edit-order
// @Accept json
// @Produce json
// @Param id path integer true "id of the order"
// @Param input body models.SwagOrder true "order update info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/orders/{id} [put]
func EditOrderByID(c *gin.Context) {
	userRole := c.GetString(userRoleCtx)
	if userRole != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.EditOrderByID] error editing order: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	var orderInput models.Order
	if err := c.ShouldBindJSON(&orderInput); err != nil {
		logger.Error.Printf("[controllers.EditOrderByID] error editing order: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input data",
		})
		return
	}

	updatedOrder, err := service.EditOrderByID(id, orderInput)
	if err != nil {
		logger.Error.Printf("[controllers.EditOrderByID] error editing order: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedOrder)
}

// DeleteOrder
// @Summary Delete Order By ID
// @Security ApiKeyAuth
// @Tags orders
// @Description delete order by ID
// @ID delete-order-by-id
// @Param id path integer true "id of the order"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/orders/{id} [delete]
func DeleteOrderByID(c *gin.Context) {
	userRole := c.GetString(userRoleCtx)
	if userRole != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.DeleteOrderByID] error deleating order: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	err = service.DeleteOrderByID(id)
	if err != nil {
		logger.Error.Printf("[controllers.DeleteOrderByID] error deleating order: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "order deleted successfully",
	})
}

// CreateCheck
// @Summary Create Check
// @Security ApiKeyAuth
// @Tags checks
// @Description create new check
// @ID create-check
// @Accept json
// @Produce json
// @Param input body models.Check true "new check info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/checks [post]
func CreateCheck(c *gin.Context) {
	var items []models.CheckItem
	if err := c.BindJSON(&items); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	orderID, _ := strconv.Atoi(c.Param("order_id"))
	tableNumber, _ := strconv.Atoi(c.Param("table_number"))

	// Вызываем сервис для создания чека
	check, err := service.CreateCheck(orderID, tableNumber, items)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем результат
	c.JSON(http.StatusOK, gin.H{"check": check})
}

func GetAllChecks(c *gin.Context) {
	checks, err := service.GetAllChecks()
	if err != nil {
		logger.Error.Printf("[controllers.GetAllChecks] error getting all checks: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"checks": checks,
	})
}

func GetCheckByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetCheckByID] error getting check %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})

		return
	}

	check, err := service.GetCheckByID(id)
	if err != nil {
		logger.Error.Printf("[controllers.GetCheckByID] error getting check %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, check)
}
