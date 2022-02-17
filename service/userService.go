package service

import (
	"net/http"
	"ticket_app/entity"
	"ticket_app/middleware"

	"github.com/gin-gonic/gin"
)

type NewUserRepository interface {
	// SignIn(email string, password string) (message string, err error)
	LogIn(email string, password string) gin.HandlerFunc
	// LogOut(email string) string
}

type newUserConnection struct {
	message string
}

func NewUserService(tes string) NewUserRepository {
	return &newUserConnection{
		message: tes,
	}
}

var newPassenger = entity.Passenger{}

func (userService *newUserConnection) LogIn(email string, password string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var passenger entity.Passenger

		if err := c.ShouldBindJSON(&passenger); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": "invalid json provide",
				"error":   http.StatusUnprocessableEntity,
			})
			return
		}

		if newPassenger.UserName != passenger.UserName || newPassenger.Password != passenger.Password {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid Email or Password",
				"error":   http.StatusUnauthorized,
			})
		}

		ts, err := middleware.CreateToken(newPassenger.ID)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, err.Error())
		}
	}
}
