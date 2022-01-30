package config

import (
	"fmt"
	"os"
	"ticket_app/entity"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("error load env")
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

	//Open Connection to database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("error connecting database")
	}
	//Closing the database after transaction
	dbClose, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer dbClose.Close()

	//Automigrate migrate data for given model
	db.AutoMigrate(
		&entity.Passenger{},
		&entity.Level{},
		&entity.Transportation{},
		&entity.TransportationType{},
		&entity.Officer{},
		&entity.Order{},
		&entity.Route{},
	)
	return db
}
