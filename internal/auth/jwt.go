package auth

import (
	"coachee-backend/gen/coachee"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secret = []byte("Zq4t7w!z%C*F-JaNdRgUjXn2r5u8x/A?")

func CreateUserToken(id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["scopes"] = []string{"client"}
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix() //Token expires after 30 minutes
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ParseToken(token string) (map[string]interface{}, error) {
	claims := make(jwt.MapClaims)

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (i interface{}, e error) {
		return secret, nil
	})
	if err != nil {
		return nil, coachee.MakeUnauthorized(errors.New("invalid token"))
	}

	return claims, nil
}
