package uc

import "github.com/paulwerner/bookkeeper/domain"

func (i interactor) AccountsFind(uID domain.UserID) ([]domain.Account, error) {
	accounts, err := i.accountRW.FindAll(uID)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
