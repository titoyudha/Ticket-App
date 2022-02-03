package entity

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID            uuid.UUID `gorm:"primaryKey" json:"id" validate:"required"`
	OrderCode     uuid.UUID `json:"order_code"`
	OrderDate     time.Time `json:"order_date" gorm:"type:time"`
	OrderPlace    string    `json:"order_place" gorm:"type:varchar(100)"`
	IDPassenger   Passenger `json:"id_passenger"`
	SeatCode      string    `json:"seat_code" gorm:"type:varchar(100)"`
	IDRoute       uuid.UUID `gorm:"primaryKey;not null" json:"id_route"`
	Route         Route     `json:"route"`
	Destination   string    `json:"destination" gorm:"type:varchar(100)"`
	DepartureDate time.Time `json:"departure_date" gorm:"type:time"`
	TimeCheckIn   time.Time `json:"time_check_in" gorm:"type:time"`
	TimeDepart    time.Time `json:"time_depart" gorm:"type:time"`
	PaymentPrice  int       `json:"payment_price" gorm:"type:integer"`
	IDOfficer     uuid.UUID `json:"id_officer"`
	Officer       Officer   `json:"officer"`
}
