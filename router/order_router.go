package router

import (
	"ticket_app/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var a controller.OrderRepository

func OrderRouter(router *gin.Engine, connect *gorm.DB) {
	router.POST("/order", a.CreateOrder())
}
