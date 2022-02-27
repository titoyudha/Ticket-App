package middleware

import (
	"ticket_app/constant"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateTokenJWT(userID string) (string, error) {

	//Generate Payload
	claims := jwt.StandardClaims{}
	// claims.Audience = userID
	claims.ExpiresAt = time.Now().Add(time.Hour * 1).Unix()

	//Token Header
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Token
	return token.SignedString(constant.JWT_SECRET)
}

// func GetUserIDFromJWT(c gin.Context)string{
// 	user := c.Param("user").(*jwt.Token)
// 	if user != nil{
// 		cla
// 	}
// }
