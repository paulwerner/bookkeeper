package uc

import d "github.com/paulwerner/bookkeeper/pkg/domain"

func (i interactor) AccountFind(id d.AccountID, uID d.UserID) (*d.Account, error) {
	account, err := i.accountRW.FindByIDAndUser(id, uID)
	if err != nil {
		return nil, err
	}
	return account, nil
}
