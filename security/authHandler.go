package security

import (
	d "github.com/paulwerner/bookkeeper/domain"
	"github.com/paulwerner/bookkeeper/uc"
	"github.com/paulwerner/bookkeeper/utils"
	"golang.org/x/crypto/bcrypt"
)

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

func (*authHandler) GenUserToken(id d.UserID) (string, error) {
	return utils.GenUserToken(id)
}

func (*authHandler) GetUserID(token string) (d.UserID, error) {
	return utils.GetUserIDFromToken(token)
}
