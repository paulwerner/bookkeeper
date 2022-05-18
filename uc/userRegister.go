package uc

import d "github.com/paulwerner/bookkeeper/domain"

func (i interactor) UserRegister(id d.UserID, name, password string) (user *d.User, token string, err error) {
	_, err = i.userRW.FindByName(name)
	if err != nil && err != d.ErrNotFound {
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
