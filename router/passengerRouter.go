package router

import (
	"ticket_app/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var a = controller.PassangerRepository{}

func PassengerRouter(router *gin.Engine, connect *gorm.DB) {

	router.POST("/passenger", a.CreatePassenger())
	router.GET("/passenger/:id", a.GetPassengerById())
	router.GET("/passengers", a.GetPassenger())
	router.PUT("/passengers/:id", a.UpdatePassenger())
	router.DELETE("/passenger/:id", a.DeletePassenger())
}
