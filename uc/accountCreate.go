package uc

import d "github.com/paulwerner/bookkeeper/domain"

func (self interactor) AccountCreate(cmd AccountCreateCmd) (account *d.Account, err error) {
	_, err = self.accountRW.FindByUserAndName(cmd.uID, cmd.name)
	if err != nil && err != d.ErrNotFound {
		err = d.ErrInternalError
		return
	}
	account, err = self.accountRW.Create(
		cmd.id,
		cmd.uID,
		cmd.name,
		cmd.description,
		cmd.accountType,
		cmd.currentBalanceValue,
		cmd.currentBalanceCurrency,
	)
	return
}
