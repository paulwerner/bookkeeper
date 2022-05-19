package uc

import d "github.com/paulwerner/bookkeeper/pkg/domain"

func (i interactor) TransactionFind(id d.TransactionID, aID d.AccountID) (*d.Transaction, error) {
	tx, err := i.transactionRW.FindByIDAndAccount(id, aID)
	if err != nil {
		return nil, err
	}
	return tx, nil
}
