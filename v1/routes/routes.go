package routes

import (
	"bitslearn/v1/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/auth/register", controllers.RegisterUser)
		api.POST("/auth/login", controllers.LoginUser)
	}
}
