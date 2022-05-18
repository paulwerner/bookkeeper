package uc

import d "github.com/paulwerner/bookkeeper/domain"

func (self interactor) TransactionCreate(cmd TransactionCreateCmd) (tx *d.Transaction, err error) {
	if tx, err = self.transactionRW.Create(
		cmd.id,
		cmd.aID,
		cmd.description,
		cmd.amountValue,
		cmd.amountCurrency,
	); err != nil {
		err = d.ErrInternalError
	}
	return
}
