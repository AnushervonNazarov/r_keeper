package controllers

import (
	"net/http"
	"r_keeper/logger"
	"r_keeper/models"
	"r_keeper/pkg/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllOrders(c *gin.Context) {
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

func GetOrderByID(c *gin.Context) {
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
	}

	c.JSON(http.StatusOK, order)
}

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

func EditOrderByID(c *gin.Context) {
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

func DeleteOrderByID(c *gin.Context) {
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
