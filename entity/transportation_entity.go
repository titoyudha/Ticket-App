package entity

import "github.com/google/uuid"

type Transportation struct {
	ID                   uuid.UUID          `gorm:"primaryKey;not null" json:"id" validate:"required"`
	Code                 string             `json:"code" gorm:"type:varchar(100)"`
	TotalSeat            int                `json:"total_seat" gorm:"type:integer(100)"`
	Description          string             `json:"description" gorm:"type:varchar(100)"`
	IDTransportationType uuid.UUID          `json:"id___transportation_type"`
	TransportationType   TransportationType `json:"transportation_type"`
	TicketPrice          int
}
