package security

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	d "github.com/paulwerner/bookkeeper/domain"
	"github.com/paulwerner/bookkeeper/uc"
	"golang.org/x/crypto/bcrypt"
)

var tokenTimeToLive = 2 * time.Hour

type authHandler struct{}

func NewAuthHandler() uc.AuthHandler {
	return &authHandler{}
}

func (*authHandler) CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (*authHandler) EncryptPassword(password string) (string, error) {
	if len(password) < 8 {
		return "", d.ErrInvalidPasswordLength
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", d.ErrInternalError
	}
	return string(hash), nil
}

func (*authHandler) GenUserToken(id d.UserID) (token string, err error) {
	token, err = jwt.
		NewWithClaims(jwt.SigningMethodHS256, newUserClaims(id, tokenTimeToLive)).
		SignedString([]byte(os.Getenv("JWT_SECRET")))
	return
}

func (*authHandler) GetUserID(t string) (d.UserID, error) {
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
	jwt.StandardClaims
}

func newUserClaims(id d.UserID, ttl time.Duration) *userClaims {
	return &userClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			Issuer:    "bookkeeper-api",
		},
	}
}
