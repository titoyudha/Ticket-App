package dto

import "time"

type PassengerCreate struct {
	Name     string `json:"name" validate:"required,min=2"`
	Address  string `json:"address" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type PassengerEdit struct {
	Name            string `json:"name" validate:"required,min=2"`
	Address         string `json:"address" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	NewPassword     string `json:"new_password" validate:"required"`
	RetypePassword  string `json:"retype_password"`
	ConfirmPassword string `json:"confirm_password"`
}

type PassengerLogIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PassengerResponse struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
	DeleteAt  time.Time `json:"delete_at"`
}

type PassengerOrder struct {
	Name        string
	Address     string
	Gender      string
	PhoneNumber string
}
