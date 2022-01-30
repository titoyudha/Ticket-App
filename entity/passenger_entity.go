package entity

import "github.com/google/uuid"

type Passenger struct {
	id            uuid.UUID `json:"id,omitempty"`
	UserName      string    `json:"user_name,omitempty"`
	Password      string    `json:"password,omitempty"`
	PassengerName string    `json:"passenger_name,omitempty"`
	Address       string    `json:"address,omitempty"`
	Gender        string    `json:"gender,omitempty"`
	PhoneNumber   string    `json:"phone_number,omitempty"`
}
