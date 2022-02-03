package entity

import "github.com/google/uuid"

type Officer struct {
	ID          uuid.UUID `json:"id" gorm:"PrimaryKey" validate:"required"`
	UserName    string    `json:"user_name" gorm:"type:varchar(100)"`
	Password    string    `json:"password" gorm:"type:varchar(50)"`
	OfficerName string    `json:"officer_name" gorm:"type:varchar(100)"`
	IDLevel     uuid.UUID `json:"id_level" gorm:""`
	Levels      Level     `json:"levels" gorm:""`
}
