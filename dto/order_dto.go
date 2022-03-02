package dto

import (
	"time"
)

type CreateOrder struct {
	OrderCode      string
	OrderPlace     string
	DeparturePlace string
	Passenger      []PassengerOrder
	SeatCode       string
	Destination    string
	DepartureDate  time.Time
	CheckIn        time.Time
	DepartureTime  time.Time
}

type UpdateOrder struct {
	Passenger     []PassengerOrder
	SeatCode      string
	Destination   string
	DepartureDate string
}

type OrderWaitList struct {
	Order   []CreateOrder
	Status  string
	IsValid bool
}
