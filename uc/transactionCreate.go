package uc

import d "github.com/paulwerner/bookkeeper/domain"

func (i interactor) TransactionCreate(
	id d.TransactionID,
	aID d.AccountID,
	description *string,
	amount int64,
	currency string,
) (tx *d.Transaction, err error) {
	if tx, err = i.transactionRW.Create(
		id,
		aID,
		description,
		amount,
		currency,
	); err != nil {
		err = d.ErrInternalError
	}
	return
}
