package entity

import "github.com/google/uuid"

type TransportationType struct {
	id          uuid.UUID `json:"id,omitempty"`
	Name_Type   string    `json:"name___type,omitempty"`
	Description string    `json:"description,omitempty"`
}
