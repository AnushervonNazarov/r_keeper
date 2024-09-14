package controllers

import (
	"net/http"
	"r_keeper/configs"
	_ "r_keeper/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RunRouts() *gin.Engine {
	r := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", PingPong)

	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", SignUp)
		auth.POST("/sign-in", SignIn)
	}

	apiG := r.Group("/api", checkUserAuthentication)

	userG := apiG.Group("/users")
	{
		userG.GET("", GetAllUsers)
		userG.GET("/:id", GetUserByID)
		userG.POST("", CreateUser)
		userG.PUT("/:id", EditUserByID)
		userG.DELETE("/:id", DeleteUserByID)
	}

	orderG := apiG.Group("/orders")
	{
		orderG.GET("", GetAllOrders)
		orderG.GET("/:id", GetOrderByID)
		orderG.POST("", CreateOrder)
		orderG.PUT("/:id", EditOrderByID)
		orderG.DELETE("/:id", DeleteOrderByID)
	}

	menuG := apiG.Group("/menus")
	{
		menuG.GET("", GetAllMenus)
		menuG.GET("/:id", GetMenuByID)
		menuG.POST("", CreateMenu)
		menuG.PUT("/:id", EditMenuByID)
		menuG.DELETE("/:id", DeleteMenuByID)
	}

	tableG := apiG.Group("/tables")
	{
		tableG.GET("", GetAllTables)
		tableG.GET("/:id", GetTableByID)
		tableG.POST("", CreateTable)
		tableG.PUT("/:id", EditTableByID)
		tableG.DELETE("/:id", DeleteTableByID)
	}

	// dishG := r.Group("/dishes")
	// {
	// 	dishG.GET("", GetALlDishes)
	// 	dishG.GET("/:id", GetDishByID)
	// 	dishG.POST("", CreateDish)
	// 	dishG.PUT("/:id", EditDishByID)
	// 	dishG.DELETE("/:id", DeleteDishByID)
	// }

	return r
}

func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
