package store

import (
	"database/sql"

	d "github.com/paulwerner/bookkeeper/domain"
	"github.com/paulwerner/bookkeeper/uc"
)

type transactionStore struct {
	db *sql.DB
}

func NewtransactionStore(db *sql.DB) uc.TransactionRW {
	return &transactionStore{
		db: db,
	}
}

func (self *transactionStore) Create(id d.TransactionID, aID d.AccountID, description *string, amountValue int64, amountCurrency string) (tx *d.Transaction, err error) {
	sqlStatement := `INSERT INTO transactions (
		id,
		account_id,
		description,
		amount_value,
		amount_currency
	) VALUES ($1, $2, $3, $4, $5)`
	if _, err = self.db.Exec(sqlStatement, id, aID, description, amountValue, amountCurrency); err != nil {
		err = d.ErrInternalError
		return
	}
	tx, err = self.FindByIDAndAccount(id, aID)
	return
}

func (self *transactionStore) FindByIDAndAccount(id d.TransactionID, aID d.AccountID) (tx *d.Transaction, err error) {
	sqlStatement := `SELECT id, description, amount_value, amount_currency FROM transactions WHERE id=$1 AND account_id=$2`
	if err = self.db.QueryRow(sqlStatement, id, aID).Scan(&tx); err != nil {
		switch err {
		case sql.ErrNoRows:
			err = d.ErrNotFound
		default:
			err = d.ErrInternalError
		}
	}
	return
}
