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

// CreateMenu
// @Summary Create Menu
// @Security ApiKeyAuth
// @Tags menus
// @Description create new menu
// @ID create-menu
// @Accept json
// @Produce json
// @Param input body models.SwagMenu true "new menu info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/menus [post]
func CreateMenu(c *gin.Context) {
	userRole := c.GetString(userRoleCtx)
	if userRole != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	var menu models.Menu
	if err := c.BindJSON(&menu); err != nil {
		logger.Error.Printf("[controllers.CreateMenu] error creating menu %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	err := service.CreateMenu(menu)
	if err != nil {
		logger.Error.Printf("[controllers.CreateMenu] error creating menu %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "menu created successfully",
	})
}

// GetAllMenus
// @Summary Get All Menus
// @Security ApiKeyAuth
// @Tags menus
// @Description get list of menus
// @ID get-all-menus
// @Produce json
// @Param q query string false "fill if you need search"
// @Success 200 {array} models.SwagMenu
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/menus [get]
func GetAllMenus(c *gin.Context) {
	menus, err := service.GetAllMenus()
	if err != nil {
		logger.Error.Printf("[controllers.GetAllMenus] error getting all menus: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"menus": menus,
	})
}

// GetMenuByID
// @Summary Get Menu By ID
// @Security ApiKeyAuth
// @Tags menus
// @Description get menu by ID
// @ID get-menu-by-id
// @Produce json
// @Param id path integer true "id of the order"
// @Success 200 {object} models.SwagMenu
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/menus/{id} [get]
func GetMenuByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetMenuByID] error getting menu %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})

		return
	}

	menu, err := service.GetMenuByID(id)
	if err != nil {
		logger.Error.Printf("[controllers.GetMenuByID] error getting menu %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, menu)
}

// EditMenuByID
// @Summary Edit Menu
// @Security ApiKeyAuth
// @Tags menus
// @Description edit existed menu
// @ID edit-menu
// @Accept json
// @Produce json
// @Param id path integer true "id of the order"
// @Param input body models.SwagMenu true "menu update info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/menus/{id} [put]
func EditMenuByID(c *gin.Context) {
	userRole := c.GetString(userRoleCtx)
	if userRole != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.EditMenuByID] error editing menu: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	var menuInput models.Menu
	if err := c.ShouldBindJSON(&menuInput); err != nil {
		logger.Error.Printf("[controllers.EditMenuByID] error editing menu: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input data",
		})
		return
	}

	updatedMenu, err := service.EditMenuByID(id, menuInput)
	if err != nil {
		logger.Error.Printf("[controllers.EditMenuByID] error editing menu: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedMenu)
}

// DeleteMenu
// @Summary Delete Menu By ID
// @Security ApiKeyAuth
// @Tags menus
// @Description delete menu by ID
// @ID delete-menu-by-id
// @Param id path integer true "id of the menu"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/menus/{id} [delete]
func DeleteMenuByID(c *gin.Context) {
	userRole := c.GetString(userRoleCtx)
	if userRole != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.DeleteMenuByID] error deleating menu: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	err = service.DeleteMenuByID(id)
	if err != nil {
		logger.Error.Printf("[controllers.DeleteMenuByID] error deleating menu: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "menu deleted successfully",
	})
}
