package service

import (
	"ticket_app/config"
	"ticket_app/dto"
	"ticket_app/entity"
	"ticket_app/helper"

	"github.com/google/uuid"
)

func Register(userCreate dto.PassengerCreate) (entity.User, error) {
	hash, _ := helper.HashPassword(userCreate.Password)
	var (
		passenger entity.User
	)

	passenger.ID = uuid.New().String()
	passenger.UserFullName = userCreate.Name
	passenger.Email = userCreate.Email
	passenger.Password = hash

	err := config.ConnectDB().Create(&passenger).Error
	if err != nil {
		return passenger, err
	}
	return passenger, nil
}

func LogIn(userLogin dto.PassengerLogIn) (entity.User, error) {
	var passengerDB entity.User

	err := config.ConnectDB().Where("email", userLogin.Email).First(&passengerDB).Error
	checkHash, _ := helper.CheckHashPassword(userLogin.Password, passengerDB.Password)

	if err != nil && !checkHash {
		return passengerDB, err
	}

	return passengerDB, nil
}

func GetAllUserData() (dataResult []entity.User, err error) {
	err = config.ConnectDB().Find(&dataResult).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetUserDetail(userID string) (entity.User, error) {
	var passengerDB entity.User

	err := config.ConnectDB().First(&passengerDB, userID).Error
	if err != nil {
		return passengerDB, err
	}
	return passengerDB, nil
}

func CheckHashPassword(password, userID string) (verified bool, err error) {
	var passenger entity.User

	err = config.ConnectDB().Where("id = ?", userID).First(&passenger).Error
	confirmedUser, _ := helper.CheckHashPassword(password, passenger.Password)

	if !confirmedUser {
		return false, err
	}

	return true, err
}

func EditUser(userToEdit dto.PassengerEdit, userID string) (entity.User, error) {
	hash, _ := helper.HashPassword(userToEdit.NewPassword)
	var passenger entity.User

	err := config.ConnectDB().First(&passenger, userID).Error

	passenger.UserFullName = userToEdit.Name
	passenger.Email = userToEdit.Email
	passenger.Password = hash

	config.ConnectDB().Save(&passenger)

	if err != nil {
		return passenger, err
	}

	return passenger, nil
}

func DeleteUser(userID string) (entity.User, error) {
	var passenger entity.User

	err := config.ConnectDB().Where("id = ?", userID).Delete(&passenger).Error

	if err != nil {
		return passenger, err
	}
	return passenger, nil
}
