package entity

import "github.com/google/uuid"

type Route struct {
	id               uuid.UUID      `json:"id,omitempty"`
	Destination      string         `json:"destination,omitempty"`
	StartingRoute    string         `json:"starting_route,omitempty"`
	FinalRoute       string         `json:"final_route,omitempty"`
	Price            int            `json:"price,omitempty"`
	IDTransportation Transportation `json:"id_transportation,omitempty"`
}
