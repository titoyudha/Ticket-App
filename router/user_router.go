package router

import (
	"ticket_app/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var userRepository controller.UserConnection

func UserRouter(router *gin.Engine, connect *gorm.DB) {
	router.POST("user", userRepository.CreateUser())
}
