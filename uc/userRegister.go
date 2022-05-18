package uc

import d "github.com/paulwerner/bookkeeper/domain"

func (self interactor) UserRegister(cmd UserRegisterCmd) (user *d.User, token string, err error) {
	_, err = self.userRW.FindByName(cmd.name)
	if err != nil && err != d.ErrNotFound {
		return
	}
	encryptedPassword := self.authHandler.EncryptPassword(cmd.password)
	user, err = self.userRW.Create(cmd.id, cmd.name, encryptedPassword)
	if err != nil {
		return
	}
	token, err = self.authHandler.GenUserToken(cmd.name)
	return
}
