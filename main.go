package main

import (
	"bitslearn/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// database init
	config.InitialDB()

	// env load
	config.LoadEnv()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "welcome to our bits learn app!",
		})
	})

	// log
	fmt.Println("🚀 Server running at http://localhost/3000")

	r.Run(":3000")
}
