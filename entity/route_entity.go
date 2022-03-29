package entity

import "github.com/google/uuid"

type Route struct {
	IDRoute        uuid.UUID      `gorm:"primaryKey;not null" json:"id,omitempty" validate:"required"`
	StartingPlace  string         `json:"starting_route,omitempty" gorm:"type:varchar(100)"`
	FinalPlace     string         `json:"final_route,omitempty" gorm:"type:varchar(100)"`
	Price          int            `json:"price,omitempty" gorm:"type:integer"`
	Transportation Transportation `json:"transportation,omitempty"`
}
