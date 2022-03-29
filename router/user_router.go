package router

import (
	"ticket_app/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRouter(router *gin.Engine, connect *gorm.DB) {
	router.POST("/user", controller.CreateUser())
}
