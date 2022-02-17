package entity

import (
	"time"
)

type Order struct {
	ID             string    `gorm:"primaryKey;unique;type:varchar(100)" json:"id" validate:"required"`
	OrderCode      string    `json:"order_code" gorm:"type:varchar(100)"`
	OrderDate      time.Time `json:"order_date" gorm:"type:time"`
	DeparturePlace string    `json:"departure_place" gorm:"type:varchar(100)"`
	// IDPassenger string      `gorm:"not null;type:varchar(100)" json:"id_passenger"`
	Passenger   []Passenger `gorm:"many2many:passenger_order" json:"passenger"`
	userIdBuyer string      `gorm:"type:varchar(100)" json:"buyer"`
	SeatCode    string      `json:"seat_code" gorm:"type:varchar(100)"`
	// IDRoute       uuid.UUID `gorm:"primaryKey;not null" json:"id_route"`
	// Route         Route     `json:"route"`
	Destination    string         `json:"destination" gorm:"type:varchar(100)"`
	DepartureDate  time.Time      `json:"departure_date" gorm:"type:time"`
	TimeCheckIn    time.Time      `json:"time_check_in" gorm:"type:time"`
	TimeDepart     time.Time      `json:"time_depart" gorm:"type:time"`
	TotalPrice     uint           `json:"payment_price" gorm:"type:integer"`
	Transportation Transportation ``
	// IDOfficer     uuid.UUID `json:"id_officer"`
	// Officer       Officer   `json:"officer"`
}
