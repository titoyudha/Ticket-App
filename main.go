package main

import (
	"ticket_app/config"
	"ticket_app/router"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	connect := config.ConnectDB()

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "server is running",
		})
	})
	router.PassengerRouter(server, connect)
	router.OrderRouter(server, connect)

	server.Run()
}
