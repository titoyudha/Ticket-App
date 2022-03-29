package entity

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	IDOrder        uuid.UUID        `json:"id_order,omitempty"`
	OrderDate      time.Time        `json:"order_date,omitempty"`
	User           User             `json:"user" gorm:"foreignkey;user_id"`
	Transportation []Transportation `json:"transportation,omitempty"`
	Passenger      []Passenger      `json:"passenger,omitempty"`
	Payment        []Payment        `json:"payment,omitempty"`
	DeparturePlace string           `json:"departure_place,omitempty"`
	TicketAmount   int              `json:"ticket_amount,omitempty"`
	TotalPrice     int              `json:"total_price,omitempty"`
}
