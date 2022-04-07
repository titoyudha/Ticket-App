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

func (u UserConnection) GetUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user entity.User
		db := config.ConnectDB()

		if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "user with that id not found",
			})
			return
		}

		result := u.GetUserByID()
		c.JSON(http.StatusOK, gin.H{
			"message": "Succes find user",
			"data":    result,
		})

		config.CloseDB(*db)
	}
}

func (u UserConnection) GetAllUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user []entity.User
		db := config.ConnectDB()

		if err := db.Find(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "data not found",
				"data":    nil,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "succes get all user data",
			"data":    &user,
		})
		config.CloseDB(*db)
	}
}

func (u UserConnection) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user entity.User
		db := config.ConnectDB()
		config.CloseDB(*db)
		param := c.Param("id")

		if err := db.Where("id = ?", param).First(&user).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "User with that id not found",
			})
			return
		}

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": "Failed Update User",
			})
			return
		}

		result := db.Save(&user)
		if result.RowsAffected <= 0 {
			panic("Error inserting to db")
		}
		c.JSON(http.StatusCreated, gin.H{
			"message": "Succes Updating User",
			"data":    user,
		})
	}

}

func (u UserConnection) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user entity.User
		db := config.ConnectDB()
		config.CloseDB(*db)

		param := c.Param("id")

		if err := db.Where("id = ?", param).First(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "User not found",
			})
			return
		}

		result := db.Delete(&user)
		if result.RowsAffected <= 0 {
			panic("Error")
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Succesfull Delete User",
		})

	}
}
