package controller

import (
	"net/http"
	"ticket_app/config"
	"ticket_app/entity"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// var db = config.ConnectDB()

type PassangerRepository struct {
	db *gorm.DB
}

func newPassangerRepository(db *gorm.DB) *PassangerRepository {
	return &PassangerRepository{db: db}
}

func (r PassangerRepository) CreatePassenger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// dsn := "root:Titoyudha*111293@tcp(127.0.0.1:3306)/ticket_app?charset=utf8mb4&parseTime=True&loc=Local"
		// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		// if err != nil {
		// 	panic(err.Error())
		// }

		var passangers = entity.Passenger{}
		// var pRepo PassangerRepository
		var db = config.ConnectDB()

		sql, err := db.DB()
		if err != nil {
			panic(err)
		}

		defer sql.Close()

		if err := c.BindJSON(&passangers); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
			return
		}

		passangers.ID = uuid.NewString()

		result := db.Create(&passangers)

		if result.RowsAffected <= 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error when insert to db",
				"data":    "",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "succes",
			"data":    passangers,
		})

	}

}

func (r *PassangerRepository) Create2(connect *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		err := connect.Create(&entity.Passenger{
			ID:            uuid.New().String(),
			UserName:      "test",
			Password:      "tes",
			PassengerName: "tes",
			Address:       "tes",
			Gender:        "",
			PhoneNumber:   "",
		})
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "succes",
			"data":    "result",
		})
	}
}
