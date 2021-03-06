package store

import (
	"database/sql"

	d "github.com/paulwerner/bookkeeper/pkg/domain"
	"github.com/paulwerner/bookkeeper/pkg/uc"
)

type transactionStore struct {
	db *sql.DB
}

func NewTransactionStore(db *sql.DB) uc.TransactionStore {
	return &transactionStore{
		db: db,
	}
}

func (ts *transactionStore) Create(tx d.Transaction) (*d.Transaction, error) {
	sqlStatement := `INSERT INTO transactions (
		id,
		account_id,
		description,
		amount,
		currency
	) VALUES ($1, $2, $3, $4, $5)`
	if _, err := ts.db.Exec(sqlStatement, tx.ID, tx.Account.ID, tx.Description, tx.Amount, tx.Currency); err != nil {
		return nil, err
	}
	return ts.FindByIDAndAccount(tx.ID, tx.Account.ID)
}

func (ts *transactionStore) FindAll(aID d.AccountID) (txs []d.Transaction, err error) {
	sqlStatement := `SELECT id, account_id, description, amount, currency FROM transactions WHERE account_id=$1`
	rows, err := ts.db.Query(sqlStatement, aID)
	defer rows.Close()
	if err != nil && err != sql.ErrNoRows {
		err = d.ErrInternalError
		return
	}
	for rows.Next() {
		var tx d.Transaction
		err = rows.Scan(&tx.ID, &tx.Account.ID, &tx.Description, &tx.Amount, &tx.Currency)
		if err != nil {
			err = d.ErrInternalError
			return
		}
		txs = append(txs, tx)
	}
	return
}

func (ts *transactionStore) FindByIDAndAccount(id d.TransactionID, aID d.AccountID) (*d.Transaction, error) {
	var tx d.Transaction
	sqlStatement := `SELECT id, account_id, description, amount, currency FROM transactions WHERE id=$1 AND account_id=$2`
	if err := ts.db.QueryRow(sqlStatement, id, aID).
		Scan(&tx.ID, &tx.Account.ID, &tx.Description, &tx.Amount, &tx.Currency); err != nil {
		switch err {
		case sql.ErrNoRows:
			err = d.ErrNotFound
		default:
			err = d.ErrInternalError
		}
		return nil, err
	}
	return &tx, nil
}
