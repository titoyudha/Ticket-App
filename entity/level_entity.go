package entity

import "github.com/google/uuid"

type Level struct {
	id   uuid.UUID `json:"id,omitempty"`
	Name string    `json:"name,omitempty"`
}
