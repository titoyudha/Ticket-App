package router

import (
	"ticket_app/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var userRepository controller.UserConnection

func UserRouter(router *gin.Engine, connect *gorm.DB) {
	UserRouter := gin.Default()
	UserRouter.Group("/api/order")
	{
		UserRouter.POST("/", userRepository.CreateUser())
		UserRouter.GET("/", userRepository.GetAllUser())
		UserRouter.GET("/:id", userRepository.GetUserByID())
		UserRouter.PUT("/:id", userRepository.UpdateUser())
		UserRouter.DELETE("/:id", userRepository.DeleteUser())
	}
}
