package uc

import d "github.com/paulwerner/bookkeeper/domain"

func (self interactor) UserRegister(
	id d.UserID,
	name string,
	password string,
) (*d.User, string, error) {
	_, err := self.userRW.FindByName(name)
	if err != nil && err != d.ErrNotFound {
		return nil, "", err
	}
	encryptedPassword := self.authHandler.EncryptPassword(password)
	user, err := self.userRW.Create(id, name, encryptedPassword)
	if err != nil {
		return nil, "", err
	}
	token, err := self.authHandler.GenUserToken(name)
	if err != nil {
		return nil, "", err
	}
	return user, token, err
}
