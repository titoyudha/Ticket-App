package main

import (
	"net/http"
	"ticket_app/config"
	"ticket_app/router"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	dbConn := config.ConnectDB()

	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	router.OrderRouter(server, dbConn)
	router.UserRouter(server, dbConn)

	server.Run(":8082")
}
