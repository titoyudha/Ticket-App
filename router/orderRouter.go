package router

import (
	"ticket_app/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func OrderRouter(router *gin.Engine, connect *gorm.DB) {
	router.POST("/order", controller.CreateOrder())
	router.GET("/order/:id", controller.GetOrderByID())
	router.GET("/order", controller.GetAllOrder())
	router.PUT("/order/:id", controller.UpdateOrder())
	router.DELETE("/order/:id", controller.DeleteOrder())
}
