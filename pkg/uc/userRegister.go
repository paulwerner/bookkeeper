package uc

import d "github.com/paulwerner/bookkeeper/pkg/domain"

func (i interactor) UserRegister(id d.UserID, name, password string) (user *d.User, token string, err error) {
	if name == "" {
		err = d.ErrInvalidEntity
		return
	}

	foundUser, err := i.userRW.FindByName(name)
	if err != nil && err != d.ErrNotFound {
		return
	}
	if foundUser != nil {
		err = d.ErrAlreadyInUse
		return
	}

	encryptedPassword, err := i.authHandler.EncryptPassword(password)
	if err != nil {
		return
	}

	user, err = i.userRW.Create(id, name, encryptedPassword)
	if err != nil {
		return
	}
	token, err = i.authHandler.GenUserToken(id)
	return
}
