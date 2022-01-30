package entity

import "github.com/google/uuid"

type Transportation struct {
	id                   uuid.UUID          `json:"id,omitempty"`
	Code                 string             `json:"code,omitempty"`
	TotalSeat            int                `json:"total_seat,omitempty"`
	Description          string             `json:"description,omitempty"`
	IDTransportationType TransportationType `json:"id___transportation_type,omitempty"`
}
