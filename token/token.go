package token

import (
	"time"
	"github.com/golang-jwt/jwt"
)


func Createtoken(id uint) (string,error){
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	return token.SignedString([]byte("something"))
}
