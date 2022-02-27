package controller

import (
	"net/http"
	"ticket_app/config"
	"ticket_app/dto"
	"ticket_app/entity"
	"ticket_app/middleware"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PassangerRepository struct {
	db *gorm.DB
}

func (r PassangerRepository) CreatePassenger() gin.HandlerFunc {
	return func(c *gin.Context) {

		var passangers = entity.Passenger{}
		var db = config.ConnectDB()

		sql, err := db.DB()
		if err != nil {
			panic(err)
		}

		defer sql.Close()

		if err := c.ShouldBindJSON(&passangers); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
			return
		}

		passangers.ID = uuid.NewString()

		result := db.Create(&passangers)

		if result.RowsAffected <= 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error when insert to db",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "succes",
			"data":    passangers,
		})

	}

}

//generate crud passenger

func (r PassangerRepository) GetPassenger() gin.HandlerFunc {
	return func(c *gin.Context) {
		var passangers []entity.Passenger
		var db = config.ConnectDB()

		sql, err := db.DB()
		if err != nil {
			panic(err)
		}

		defer sql.Close()

		if err := db.Find(&passangers).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "succes",
			"data":    passangers,
		})

	}

}

func (r PassangerRepository) GetPassengerById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var passangers entity.Passenger
		var db = config.ConnectDB()

		sql, err := db.DB()
		if err != nil {
			panic(err)
		}

		defer sql.Close()

		if err := db.Where("id = ?", c.Param("id")).First(&passangers).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "succes",
			"data":    passangers,
		})

	}

}

func (r PassangerRepository) UpdatePassenger() gin.HandlerFunc {
	return func(c *gin.Context) {
		var passangers entity.Passenger
		var db = config.ConnectDB()

		sql, err := db.DB()
		if err != nil {
			panic(err)
		}

		defer sql.Close()

		if err := db.Where("id = ?", c.Param("id")).First(&passangers).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
			})
			return
		}

		if err := c.ShouldBindJSON(&passangers); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
			})
			return
		}

		result := db.Save(&passangers)

		if result.RowsAffected <= 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error occured when updating passenger",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data":    passangers,
		})

	}

}

func (r PassangerRepository) DeletePassenger() gin.HandlerFunc {
	return func(c *gin.Context) {
		var passangers entity.Passenger
		var db = config.ConnectDB()

		sql, err := db.DB()
		if err != nil {
			panic(err)
		}

		defer sql.Close()

		if err := db.Where("id = ?", c.Param("id")).First(&passangers).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "data with that id not found",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "succesfull delete",
			"data":    nil,
		})
	}
}

func LogInPassenger() gin.HandlerFunc {
	return func(c *gin.Context) {
		userLogIn := dto.PassengerLogIn{}
		dbConn := config.ConnectDB()
		c.BindJSON(&userLogIn)

		var user entity.Passenger
		err := dbConn.Where("email = ?", userLogIn.Email).First(&user).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "email not found",
			})
			return
		}
		token, _ := middleware.GenerateTokenJWT(user.ID)

		var userResponse = dto.PassengerResponse{
			ID:        user.ID,
			Name:      user.UserName,
			Email:     user.PhoneNumber,
			Token:     token,
			CreatedAt: time.Time{},
			UpdateAt:  time.Time{},
			DeleteAt:  time.Time{},
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Succesful Log In",
			"Data":    userResponse,
		})
	}
}
