package uc

import d "github.com/paulwerner/bookkeeper/domain"

func (i interactor) TransactionsFind(aID d.AccountID) ([]d.Transaction, error) {
	txs, err := i.transactionRW.FindAll(aID)
	if err != nil {
		return nil, err
	}
	return txs, nil
}
