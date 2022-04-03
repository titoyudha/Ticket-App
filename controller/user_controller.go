package controller

import (
	"net/http"
	"ticket_app/config"
	"ticket_app/entity"
	"ticket_app/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserController interface {
	CreateUser() gin.HandlerFunc
	GetUserByID() gin.HandlerFunc
	GetAllUser() gin.HandlerFunc
	DeleteUser() gin.HandlerFunc
	UpdateUser() gin.HandlerFunc
}

type UserRepository struct {
	userService service.UserRepository
}

type UserConnection struct {
	db *gorm.DB
}

func (db UserConnection) CreateUser() gin.HandlerFunc {
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

func GetUserByID() {

}
