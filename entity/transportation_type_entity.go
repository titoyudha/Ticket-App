package entity

import "github.com/google/uuid"

type TransportationType struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Name_Type   string    `json:"name___type" gorm:"type:varchar(100)"`
	Description string    `json:"description" gorm:"type:varchar(100)"`
}
