package entity

type User struct {
	ID           string `gorm:"id" json:"id"`
	UserName     string `json:"user_name" gorm:"user_name" validate:"required"`
	Email        string `json:"email" gorm:"email" validate:"required"`
	Password     string `json:"password" gorm:"password"`
	UserFullName string `json:"user_full_name" gorm:"user_full_name"`
	PhoneNumber  string `json:"phone_number" gorm:"phone_number"`
}
