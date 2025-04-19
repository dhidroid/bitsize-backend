package main

import (
	"bitslearn/config"
	"bitslearn/v1/routes"
	"bitslearn/v1/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// database init
	config.InitialDB()

	services.InitUserService(config.DB)

	// env load
	config.LoadEnv()

	// routers
	routes.RegisterRoutes(r)

	// welcome router
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "welcome to our bits learn app!",
		})
	})

	// log
	fmt.Println("🚀 Server running at http://localhost/3000")

	r.Run(":3000")
}
