package entity

type Transportation struct {
	ID                 string `gorm:"primaryKey;not null" json:"id" validate:"required"`
	TotalSeat          int    `json:"total_seat" gorm:"type:integer(100)"`
	TransportationName string `json:"transportation_name,omitempty"`
	Vechile            string `json:"vechile,omitempty"`
	Notes              string `json:"notes,omitempty"`
}
