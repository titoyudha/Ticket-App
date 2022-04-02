package service

import (
	"ticket_app/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user entity.User) entity.User
	GetUserByID(id string) entity.User
	GetAllUser() []entity.User
	UpdateUser(user entity.User) entity.User
	SignIn(email string, password string) (message string, err error)
	LogIn(email string, password string) error
	LogOut(email string) string
}

type userConnection struct {
	db *gorm.DB
}

//NewUserRepository create instance UserRepository
func NewUserRepository(dbConn *gorm.DB) UserRepository {
	return &userConnection{
		db: dbConn,
	}
}

func (db *userConnection) CreateUser(user entity.User) entity.User {
	db.db.Save(&user)
	return user
}

func (db *userConnection) GetUserByID(userId string) entity.User {
	var user entity.User
	db.db.Find(&user, userId)
	return user
}

func (db *userConnection) GetAllUser() []entity.User {
	var users []entity.User
	db.db.Find(&users)
	return users
}

func (db *userConnection) UpdateUser(user entity.User) entity.User {
	db.db.Save(&user)
	db.db.Find(&user)

	return user
}

func (db *userConnection) SignIn(email, password string) (message string, err error) {
	return "hello", nil
}

func (db *userConnection) LogIn(email, password string) error {
	return nil
}

func (db *userConnection) LogOut(email string) string {
	return "test"
}
