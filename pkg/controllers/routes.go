package controllers

import "github.com/gin-gonic/gin"

func RunRouts() error {
	r := gin.Default()

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

	return nil
}
