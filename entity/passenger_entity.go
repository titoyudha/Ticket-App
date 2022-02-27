package entity

type Passenger struct {
	ID            string `gorm:"primaryKey;auto_increment;type:varchar(255)" json:"id" validate:"required"`
	UserName      string `json:"user_name" gorm:"type:varchar(100)" validate:"required"`
	Email         string `json:"email", gorm:"type:varchar(100)" validate:"required,email"`
	Password      string `json:"password" gorm:"type:varchar(50)"`
	PassengerName string `json:"passenger_name" gorm:"type:varchar(100)"`
	Address       string `json:"address" gorm:"type:varchar(100)"`
	Gender        string `json:"gender" gorm:"type:varchar(10)"`
	PhoneNumber   string `json:"phone_number" gorm:"type:varchar(100)"`
	// Orders        []Order `gorm:"many2many:passenger_order"`
}
