package uc

import d "github.com/paulwerner/bookkeeper/pkg/domain"

func (i interactor) TransactionsFind(aID d.AccountID) ([]d.Transaction, error) {
	txs, err := i.transactionStore.FindAll(aID)
	if err != nil {
		return nil, err
	}
	return txs, nil
}
