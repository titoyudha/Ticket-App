package entity

import "github.com/google/uuid"

type Officer struct {
	id          uuid.UUID `json:"id,omitempty"`
	UserName    string    `json:"user_name,omitempty"`
	Password    string    `json:"password,omitempty"`
	OfficerName string    `json:"officer_name,omitempty"`
	IDLevel     Level     `json:"id_level,omitempty"`
}
