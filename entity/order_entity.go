package entity

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	id            uuid.UUID `json:"id,omitempty"`
	OrderCode     uuid.UUID `json:"order_code,omitempty"`
	OrderDate     time.Time `json:"order_date,omitempty"`
	OrderPlace    string    `json:"order_place,omitempty"`
	IDPassenger   Passenger `json:"id_passenger,omitempty"`
	SeatCode      string    `json:"seat_code,omitempty"`
	IDRoute       Route     `json:"id_route,omitempty"`
	Destination   string    `json:"destination,omitempty"`
	DepartureDate time.Time `json:"departure_date,omitempty"`
	TimeCheckIn   time.Time `json:"time_check_in,omitempty"`
	TimeDepart    time.Time `json:"time_depart,omitempty"`
	PaymentPrice  int       `json:"payment_price,omitempty"`
	IDOfficer     Officer   `json:"id_officer,omitempty"`
}
