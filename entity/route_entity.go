package entity

import "github.com/google/uuid"

type Route struct {
	ID               uuid.UUID      `gorm:"primaryKey;not null" json:"id" validate:"required"`
	Destination      string         `json:"destination" gorm:"type:varchar(100)"`
	StartingRoute    string         `json:"starting_route" gorm:"type:varchar(100)"`
	FinalRoute       string         `json:"final_route" gorm:"type:varchar(100)"`
	Price            int            `json:"price,omitempty" gorm:"type:integer"`
	IDTransportation uuid.UUID      `json:"id_transportation"`
	Transportation   Transportation `json:"transportation"`
}
