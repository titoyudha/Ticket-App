package middleware

import (
	"os"
	"ticket_app/entity"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var (
	passenger entity.Passenger
)

type AccessDetails struct {
	AccessUuid string
	UserId     int64
}

type TokensDetail struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

func CreateToken(userid string) (td *TokensDetail, err error) {
	err = godotenv.Load()
	if err != nil {
		panic("cant load env file")
	}
	td = &TokensDetail{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.New().String()
	td.AtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = td.AccessUuid + "++" + passenger.IDPassenger

	//Create Access token
	os.Setenv(os.Getenv("KEY"), os.Getenv("VALUE"))
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userid"] = passenger.IDPassenger
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	td.AccessToken, err = at.SignedString(os.Getenv("KEY"))
	if err != nil {
		return nil, err
	}

	//Create Refresh token
	os.Setenv(os.Getenv("KEY"), os.Getenv("VALUE"))
	refreshTokensClaims := jwt.MapClaims{}
	refreshTokensClaims["refresh_uuid"] = td.RefreshUuid
	refreshTokensClaims["userid"] = passenger.IDPassenger
	refreshTokensClaims["exp"] = td.RtExpires

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokensClaims)
	td.RefreshToken, err = rt.SignedString(os.Getenv("KEY"))
	if err != nil {
		return
	}
	return td, err
}
