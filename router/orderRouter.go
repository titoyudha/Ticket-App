package router

import (
	"ticket_app/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var orderRepository = controller.OrderRepository{}

func OrderRouter(router *gin.Engine, connect *gorm.DB) {
	router.POST("/order", orderRepository.CreateOrder())
	router.GET("/order/:id", orderRepository.GetOrderByID())
	router.GET("/order", orderRepository.GetAllOrder())
	router.PUT("/order/:id", orderRepository.UpdateOrder())
	router.DELETE("/order/:id", orderRepository.DeleteOrder())
}
