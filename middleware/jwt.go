package middleware

import (
	"os"
	"ticket_app/entity"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var (
	passenger entity.Passenger
)

func CreateToken(userid string) (token string, err error) {
	err = godotenv.Load()
	if err != nil {
		panic("cant load env file")
	}

	//Create Access token
	os.Setenv(os.Getenv("KEY"), os.Getenv("VALUE"))
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userid"] = passenger.ID
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = at.SignedString(os.Getenv("KEY"))
	if err != nil {
		return "", err
	}
	return token, nil
}
