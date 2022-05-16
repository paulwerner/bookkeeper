package uc

import d "github.com/paulwerner/bookkeeper/domain"

func (self interactor) UserFind(id d.UserID) (*d.User, string, error) {
	user, err := self.userRW.FindByID(id)
	if err != nil {
		return nil, "", err
	}
	token, err := self.authHandler.GenUserToken(user.Name)
	if err != nil {
		return nil, "", err
	}
	return user, token, nil
}
