package entity

type Passenger struct {
	ID            string `gorm:"primaryKey;many2many;not null" json:"id" validate:"required"`
	UserName      string `json:"user_name" gorm:"type:varchar(100)"`
	Password      string `json:"password" gorm:"type:varchar(50)"`
	PassengerName string `json:"passenger_name" gorm:"type:varchar(100)"`
	Address       string `json:"address" gorm:"type:varchar(100)"`
	Gender        string `json:"gender" gorm:"type:varchar(10)"`
	PhoneNumber   string `json:"phone_number" gorm:"type:varchar(100)"`
}
