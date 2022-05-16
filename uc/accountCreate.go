package uc

import d "github.com/paulwerner/bookkeeper/domain"

func (self interactor) AccountCreate(
	id d.AccountID,
	uID d.UserID,
	name string,
	description *string,
	accountType d.AccountType,
	currentBalanceValue int64,
	currentBalanceCurrency string,
) (*d.Account, error) {
	_, err := self.accountRW.FindByUserAndName(uID, name)
	if err != nil && err != d.ErrNotFound {
		return nil, err
	}
	account, err := self.accountRW.Create(
		id,
		uID,
		name,
		description,
		accountType,
		currentBalanceValue,
		currentBalanceCurrency,
	)
	if err != nil {
		return nil, err
	}
	return account, nil
}
