package service

import (
	"ticket_app/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}
type Passenger struct {
	ID            uuid.UUID `gorm:"primaryKey;many2many;not null" json:"id" validate:"required"`
	UserName      string    `json:"user_name" gorm:"type:varchar(100)"`
	Password      string    `json:"password" gorm:"type:varchar(50)"`
	PassengerName string    `json:"passenger_name" gorm:"type:varchar(100)"`
	Address       string    `json:"address" gorm:"type:varchar(100)"`
	Gender        string    `json:"gender" gorm:"type:varchar(10)"`
	PhoneNumber   string    `json:"phone_number" gorm:"type:varchar(100)"`
}

var passenger = entity.Passenger{}

func (passenger *Passenger) Create(db *gorm.DB) (*Passenger, error) {
	err := db.Debug().Save(&passenger).Error
	if err != nil {
		return passenger, err
	}
	return passenger, nil

}
