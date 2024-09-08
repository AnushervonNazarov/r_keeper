package controllers

import (
	"r_keeper/configs"

	"github.com/gin-gonic/gin"
)

func RunRouts() *gin.Engine {
	r := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)

	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", SignUp)
		auth.POST("/sign-in", SignIn)
	}

	userG := r.Group("/users", checkUserAuthentication)
	{
		userG.GET("", GetAllUsers)
		userG.GET("/:id", GetUserByID)
		userG.POST("", CreateUser)
		userG.PUT("/:id", EditUserByID)
		userG.DELETE("/:id", DeleteUserByID)
	}

	orderG := r.Group("/orders", checkUserAuthentication)
	{
		orderG.GET("", GetAllOrders)
		orderG.GET("/:id", GetOrderByID)
		orderG.POST("", CreateOrder)
		orderG.PUT("/:id", EditOrderByID)
		orderG.DELETE("/:id", DeleteOrderByID)
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
