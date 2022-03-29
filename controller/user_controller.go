package controller

import (
	"net/http"
	"ticket_app/config"
	"ticket_app/entity"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := entity.User{}
		var db = config.ConnectDB()

		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    http.StatusUnprocessableEntity,
				"message": "Enter Valid Format",
			})
			return
		}

		user.ID = uuid.NewString()

		result := db.Create(&user)
		if result == nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "Failed Save User",
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"code":    http.StatusCreated,
			"message": "Succes",
			"data":    user,
		})
	}
}
