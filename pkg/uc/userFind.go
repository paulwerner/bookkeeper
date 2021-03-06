package uc

import d "github.com/paulwerner/bookkeeper/pkg/domain"

func (i interactor) UserFind(id d.UserID) (*d.User, string, error) {
	user, err := i.userStore.FindByID(id)
	if err != nil {
		return nil, "", err
	}
	token, err := i.authHandler.GenUserToken(user.ID)
	if err != nil {
		return nil, "", err
	}
	return user, token, nil
}
