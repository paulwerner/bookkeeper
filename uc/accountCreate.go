package uc

import d "github.com/paulwerner/bookkeeper/domain"

func (i interactor) AccountCreate(a d.Account) (account *d.Account, err error) {
	_, err = i.accountRW.FindByUserAndName(a.User.ID, a.Name)
	if err != nil && err != d.ErrNotFound {
		err = d.ErrInternalError
		return
	}
	account, err = i.accountRW.Create(
		a.ID,
		a.User.ID,
		a.Name,
		a.Description,
		a.Type,
		a.BalanceValue,
		a.BalanceCurrency,
	)
	return
}
