package uc

import d "github.com/paulwerner/bookkeeper/pkg/domain"

func (i interactor) TransactionCreate(tx d.Transaction) (*d.Transaction, error) {
	newTx, err := i.transactionStore.Create(tx)
	if err != nil {
		err = d.ErrInternalError
	}
	tx.UpdateAccountBalance()
	_, err = i.accountStore.Update(tx.Account)
	if err != nil {
		// err = d.ErrInternalError
	}
	return newTx, err
}
