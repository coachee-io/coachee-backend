package auth

import (
	"coachee-backend/gen/coachee"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/dgrijalva/jwt-go"
)

var secret = []byte("Zq4t7w!z%C*F-JaNdRgUjXn2r5u8x/A?")

func CreateUserToken(id uint, created time.Time) (string, error) {
	claims := jwt.MapClaims{}
	claims["scopes"] = []string{"client"}
	claims["id"] = id
	claims["iat"] = created.Unix()
	claims["exp"] = created.Add(60 * time.Minute).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func CreateAdminToken(created time.Time) (string, error) {
	claims := jwt.MapClaims{}
	claims["scopes"] = []string{"admin"}
	claims["iat"] = created.Unix()
	claims["exp"] = created.Add(60 * time.Minute).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ParseToken(token string) (jwt.MapClaims, error) {
	if len(strings.Split(token, " ")) == 2 {
		token = strings.Split(token, " ")[1]
	}
	claims := make(jwt.MapClaims)

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (i interface{}, e error) {
		return secret, nil
	})
	if err != nil {
		return nil, coachee.MakeUnauthorized(errors.New("invalid token"))
	}

	return claims, nil
}

// GetCustomerID returns the id from the jwt
func GetCustomerID(token string) (uint, error) {
	if len(strings.Split(token, " ")) == 2 {
		token = strings.Split(token, " ")[1]
	}
	claims := make(jwt.MapClaims)

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (i interface{}, e error) {
		return secret, nil
	})
	if err != nil {
		log.Err(err).Msg("the error")
		return 0, coachee.MakeUnauthorized(errors.New("invalid token"))
	}

	inter, ok := claims["id"]
	if !ok {
		return 0, coachee.MakeInternal(errors.New("missing id"))
	}
	uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", inter), 10, 32)
	if err != nil {
		return 0, coachee.MakeInternal(err)
	}

	return uint(uid), nil
}
