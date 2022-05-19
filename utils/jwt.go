package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	d "github.com/paulwerner/bookkeeper/pkg/domain"
)

var tokenTimeToLive = 2 * time.Hour

func GenUserToken(id d.UserID) (token string, err error) {
	token, err = jwt.
		NewWithClaims(jwt.SigningMethodHS256, newUserClaims(id, tokenTimeToLive)).
		SignedString([]byte(os.Getenv("JWT_SECRET")))
	return
}

func GetUserIDFromToken(t string) (d.UserID, error) {
	token, err := jwt.ParseWithClaims(
		t,
		&userClaims{},
		func(token *jwt.Token) (any, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		},
	)
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*userClaims); ok && token.Valid {
		return claims.ID, nil
	}
	return "", d.ErrInternalError
}

type userClaims struct {
	ID d.UserID
	jwt.RegisteredClaims
}

func newUserClaims(id d.UserID, ttl time.Duration) *userClaims {
	return &userClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			Issuer:    "bookkeeper-api",
		},
	}
}
