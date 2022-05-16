package uc

import d "github.com/paulwerner/bookkeeper/domain"

func (self interactor) AccountFind(id d.AccountID, uID d.UserID) (*d.Account, error) {
	account, err := self.accountRW.FindByIDAndUser(id, uID)
	if err != nil {
		return nil, err
	}
	return account, nil
}
