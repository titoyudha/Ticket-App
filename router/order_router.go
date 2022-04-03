package router

import (
	"ticket_app/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var controllerRepo controller.OrderRepository

func OrderRouter(router *gin.Engine, connect *gorm.DB) {

	orderRoute := gin.Default()
	orderRoute.Group("/api/order")
	{
		orderRoute.POST("/", controllerRepo.CreateOrder())
		orderRoute.GET("/", controllerRepo.GetAllOrder())
		orderRoute.GET("/:id", controllerRepo.GetOrderByID())
		orderRoute.PUT("/:id", controllerRepo.UpdateOrder())
		orderRoute.DELETE("/:id", controllerRepo.DeleteOrder())
	}

}
