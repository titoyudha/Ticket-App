package entity

import "time"

type Payment struct {
	ID               string           `json:"id"`
	DeadlinePayment  time.Time        `json:"deadline_payment"`
	Status           string           `json:"status"`
	TotalPayment     int              `json:"total_payment"`
	PaymentMethod    string           `json:"payment_method"`
	TransportationID []Transportation `json:"transportation_id"`
	UserID           []Passenger      `json:"user_id"`
	OrderID          []Order          `json:"order_id"`
	CreatedAt        time.Time        `json:"created_at"`
	UpdatedAt        time.Time        `json:"updated_at"`
}
