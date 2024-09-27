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

// CreateTable
// @Summary Create Table
// @Security ApiKeyAuth
// @Tags tables
// @Description create new table
// @ID create-table
// @Accept json
// @Produce json
// @Param input body models.SwagTable true "new table info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/tables [post]
func CreateTable(c *gin.Context) {
	userRole := c.GetString(userRoleCtx)
	if userRole != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	var table models.Table
	if err := c.BindJSON(&table); err != nil {
		logger.Error.Printf("[controllers.CreateTable] error creating table %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	err := service.CreateTable(table)
	if err != nil {
		logger.Error.Printf("[controllers.CreateTable] error creating table %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "table created successfully",
	})
}

// GetAllTables
// @Summary Get All Tables
// @Security ApiKeyAuth
// @Tags tables
// @Description get list of tables
// @ID get-all-tables
// @Produce json
// @Param q query string false "fill if you need search"
// @Success 200 {array} models.SwagTable
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/tables [get]
func GetAllTables(c *gin.Context) {
	tables, err := service.GetAllTables()
	if err != nil {
		logger.Error.Printf("[controllers.GetAllTables] error getting all tables: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tables": tables,
	})
}

// GetTableByID
// @Summary Get Table By ID
// @Security ApiKeyAuth
// @Tags tables
// @Description get table by ID
// @ID get-table-by-id
// @Produce json
// @Param id path integer true "id of the table"
// @Success 200 {object} models.SwagTable
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/tables/{id} [get]
func GetTableByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetTableByID] error getting table %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})

		return
	}

	table, err := service.GetTableByID(id)
	if err != nil {
		logger.Error.Printf("[controllers.GetTableByID] error getting table %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, table)
}

// EdiTtableByID
// @Summary Edit Table
// @Security ApiKeyAuth
// @Tags tables
// @Description edit existed table
// @ID edit-table
// @Accept json
// @Produce json
// @Param id path integer true "id of the table"
// @Param input body models.SwagTable true "table update info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/tables/{id} [put]
func EditTableByID(c *gin.Context) {
	userRole := c.GetString(userRoleCtx)
	if userRole != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.EditTableByID] error editing table: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	var tableInput models.Table
	if err := c.ShouldBindJSON(&tableInput); err != nil {
		logger.Error.Printf("[controllers.EditTableByID] error editing table: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input data",
		})
		return
	}

	updatedTable, err := service.EditTableByID(id, tableInput)
	if err != nil {
		logger.Error.Printf("[controllers.EdittableByID] error editing table: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedTable)
}

// DeleteTableByID
// @Summary Delete Table By ID
// @Security ApiKeyAuth
// @Tags tables
// @Description delete table by ID
// @ID delete-table-by-id
// @Param id path integer true "id of the table"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/tables/{id} [delete]
func DeleteTableByID(c *gin.Context) {
	userRole := c.GetString(userRoleCtx)
	if userRole != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.DeleteTableByID] error deleating table: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	err = service.DeleteTableByID(id)
	if err != nil {
		logger.Error.Printf("[controllers.DeleteTableByID] error deleating table: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "table deleted successfully",
	})
}
