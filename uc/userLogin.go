package uc

import d "github.com/paulwerner/bookkeeper/domain"

func (i interactor) UserLogin(name, password string) (user *d.User, token string, err error) {
	user, err = i.userRW.FindByName(name)
	if err != nil {
		return
	}
	if ok := i.authHandler.CheckPassword(user.Password, password); !ok {
		err = d.ErrInvalidPassword
		return
	}
	token, err = i.authHandler.GenUserToken(user.ID)
	return
}
