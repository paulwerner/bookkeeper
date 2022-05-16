package uc

import d "github.com/paulwerner/bookkeeper/domain"

func (self interactor) TransactionCreate(
	id d.TransactionID,
	aID d.AccountID,
	description *string,
	amountValue int64,
	amountCurrency string,
) (*d.Transaction, error) {
	tx, err := self.transactionRW.Create(
		id,
		aID,
		description,
		amountValue,
		amountCurrency,
	)
	if err != nil {
		return nil, err
	}
	return tx, nil
}
