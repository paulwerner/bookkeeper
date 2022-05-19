package uc

import d "github.com/paulwerner/bookkeeper/pkg/domain"

func (i interactor) TransactionCreate(tx d.Transaction) (*d.Transaction, error) {
	newTx, err := i.transactionRW.Create(
		tx.ID,
		tx.Account.ID,
		tx.Description,
		tx.Amount,
		tx.Currency,
	)
	if err != nil {
		err = d.ErrInternalError
	}
	return newTx, err
}
