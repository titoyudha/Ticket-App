package controller

import (
	"net/http"
	"ticket_app/config"
	"ticket_app/entity"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func (r *OrderRepository) CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			order     entity.Order
			passenger entity.Passenger
			db        = config.ConnectDB()
		)

		errOrder := c.ShouldBindJSON(&order)
		if errOrder != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error 1",
			})
			return
		}

		passengerToInsert := c.ShouldBindJSON(&passenger)
		if passengerToInsert == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error 2",
			})
			return
		}
		// order.OrderDate, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		// order.TimeCheckIn, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		// order.TimeDepart, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		// order.DepartureDate, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		// order.ID = uuid.New().String()
		// order.OrderCode = uuid.New().String()
		// order.IDPassenger = uuid.New().String()

		passenger.ID = uuid.New().String()
		order.ID = uuid.New().String()
		order.OrderDate, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.TimeCheckIn, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.TimeDepart, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.DepartureDate, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.Passenger = append(order.Passenger, passenger)

		db.Create(&passenger)
		db.Create(&order)
		// if result.RowsAffected <= 0 {
		// 	panic(result)
		// }

		// if err := db.Model(neworders).Preload("Passenger").Find(&newpassengers).Error; err != nil {
		// 	panic(err)
		// }
		// db.Model(order).Association("Passenger").Append(order.Passenger)

		// for _, userOrder := range neworders.Passenger {
		// 	fmt.Println(userOrder.ID)
		// }

		c.JSON(http.StatusOK, order)

	}
}

//generate crud order
