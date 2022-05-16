package uc

import d "github.com/paulwerner/bookkeeper/domain"

func (self interactor) TransactionFind(id d.TransactionID, aID d.AccountID) (*d.Transaction, error) {
	tx, err := self.transactionRW.FindByIDAndAccount(id, aID)
	if err != nil {
		return nil, err
	}
	return tx, nil
}
