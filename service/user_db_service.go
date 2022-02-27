package service

import (
	"ticket_app/config"
	"ticket_app/dto"
	"ticket_app/entity"
	"ticket_app/helper"
)

func Register(userCreate dto.PassengerCreate) (entity.Passenger, error) {
	hash, _ := helper.HashPassword(userCreate.Password)
	var passenger entity.Passenger

	passenger.UserName = userCreate.Name
	passenger.Email = userCreate.Email
	passenger.Password = hash
	passenger.Address = userCreate.Address

	err := config.ConnectDB().Create(&passenger).Error
	if err != nil {
		return passenger, err
	}
	return passenger, nil
}

func LogIn(userLogin dto.PassengerLogIn) (entity.Passenger, error) {
	var passengerDB entity.Passenger

	err := config.ConnectDB().Where("email", userLogin.Email).First(&passengerDB).Error
	checkHash, _ := helper.CheckHashPassword(userLogin.Password, passengerDB.Password)

	if err != nil && !checkHash {
		return passengerDB, err
	}

	return passengerDB, nil
}

func GetAllUserData() (dataResult []entity.Passenger, err error) {
	err = config.ConnectDB().Find(&dataResult).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetUserDetail(userID string) (entity.Passenger, error) {
	var passengerDB entity.Passenger

	err := config.ConnectDB().First(&passengerDB, userID).Error
	if err != nil {
		return passengerDB, err
	}
	return passengerDB, nil
}
