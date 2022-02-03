package entity

import (
	"github.com/google/uuid"
)

type Level struct {
	ID   uuid.UUID `gorm:"primaryKey;not null" json:"id" validate:"required,uuid"`
	Name string    `json:"name" gorm:"type:varchar(100);not null"`
}
