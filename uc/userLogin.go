package uc

import d "github.com/paulwerner/bookkeeper/domain"

func (self interactor) UserLogin(cmd UserLoginCmd) (user *d.User, token string, err error) {
	user, err = self.userRW.FindByName(cmd.name)
	if err != nil {
		return
	}
	if ok := self.authHandler.CheckPassword(cmd.password, user.Password); !ok {
		err = d.ErrInvalidPassword
		return
	}
	token, err = self.authHandler.GenUserToken(user.Name)
	return
}
