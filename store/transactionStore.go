package store

import (
	"database/sql"

	d "github.com/paulwerner/bookkeeper/domain"
	"github.com/paulwerner/bookkeeper/uc"
)

type transactionStore struct {
	db *sql.DB
}

func NewTransactionStore(db *sql.DB) uc.TransactionRW {
	return &transactionStore{
		db: db,
	}
}

func (self *transactionStore) Create(
	id d.TransactionID,
	aID d.AccountID,
	description *string,
	amount int64,
	currency string,
) (*d.Transaction, error) {
	sqlStatement := `INSERT INTO transactions (
		id,
		account_id,
		description,
		amount,
		currency
	) VALUES ($1, $2, $3, $4, $5)`
	if _, err := self.db.Exec(sqlStatement, id, aID, description, amount, currency); err != nil {
		return nil, err
	}
	return self.FindByIDAndAccount(id, aID)
}

func (self *transactionStore) FindByIDAndAccount(id d.TransactionID, aID d.AccountID) (*d.Transaction, error) {
	var tx d.Transaction
	sqlStatement := `SELECT id, description, amount, currency FROM transactions WHERE id=$1 AND account_id=$2`
	if err := self.db.QueryRow(sqlStatement, id, aID).Scan(&tx.ID, &tx.Description, &tx.Amount, &tx.Currency); err != nil {
		switch err {
		case sql.ErrNoRows:
			err = d.ErrNotFound
		default:
			// err = d.ErrInternalError
		}
		return nil, err
	}
	return &tx, nil
}
