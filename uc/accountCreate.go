package uc

import d "github.com/paulwerner/bookkeeper/domain"

func (i interactor) AccountCreate(
	id d.AccountID,
	uID d.UserID,
	name string,
	description *string,
	accountType d.AccountType,
	currentBalanceValue int64,
	currentBalanceCurrency string,
) (account *d.Account, err error) {
	_, err = i.accountRW.FindByUserAndName(uID, name)
	if err != nil && err != d.ErrNotFound {
		err = d.ErrInternalError
		return
	}
	account, err = i.accountRW.Create(
		id,
		uID,
		name,
		description,
		accountType,
		currentBalanceValue,
		currentBalanceCurrency,
	)
	return
}
