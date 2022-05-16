package uc

import d "github.com/paulwerner/bookkeeper/domain"

func (self interactor) UserLogin(name, password string) (*d.User, string, error) {
	user, err := self.userRW.FindByName(name)
	if err != nil {
		return nil, "", err
	}
	if ok := self.authHandler.CheckPassword(password, user.Password); !ok {
		return nil, "", d.ErrInvalidPassword
	}
	token, err := self.authHandler.GenUserToken(user.Name)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
