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
		// order.Passenger = append(order.Passenger, passenger)

		// db.Create(&passenger)
		db.Create(&order)
		// if result.RowsAffected <= 0 {
		// 	panic(result)
		// }

		// if err := db.Model(neworders).Preload("Passenger").Find(&newpassengers).Error; err != nil {
		// 	panic(err)
		// }
		db.Model(order).Association("passenger_order").Append(order.Passenger)

		// for _, userOrder := range neworders.Passenger {
		// 	fmt.Println(userOrder.ID)
		// }

		c.JSON(http.StatusOK, order)

	}
}

func (r *OrderRepository) GetOrderByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var order entity.Order
		var db = config.ConnectDB()

		sql, err := db.DB()
		if err != nil {
			panic("cant connect to database")
		}
		defer sql.Close()

		if err := db.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "cant find order by that id",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "succes find order",
			"data":    &order,
		})
	}
}

func (r *OrderRepository) GetAllOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var order []entity.Order
		var db = config.ConnectDB()

		sql, err := db.DB()
		if err != nil {
			panic("error connecting to db")
		}
		defer sql.Close()

		if err := db.Find(&order).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "data not found",
				"data":    nil,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "success get all order",
			"data":    &order,
		})

	}
}

func (r *OrderRepository) UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var order entity.Order
		var db = config.ConnectDB()

		sql, err := db.DB()
		if err != nil {
			panic(err)
		}
		defer sql.Close()

		if err := db.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Order Not Found",
			})
			return
		}
		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error while updating order",
				"data":    nil,
			})
		}
		result := db.Save(&order)

		if result.RowsAffected <= 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to insert data to the database",
				"data":    nil,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "succes updating order",
			"data":    order,
		})

	}
}
func (r *OrderRepository) DeleteOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var order entity.Order
		var db = config.ConnectDB()

		sql, err := db.DB()
		if err != nil {
			panic(err)
		}
		defer sql.Close()

		if err := db.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Order Not Found",
			})
			return
		}
		db.Delete(&order)

		c.JSON(http.StatusOK, gin.H{
			"message": "succes delete order",
			"data":    order,
		})

	}
}
