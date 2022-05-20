package store

import (
	"database/sql"

	d "github.com/paulwerner/bookkeeper/pkg/domain"
	"github.com/paulwerner/bookkeeper/pkg/uc"
)

type accountStore struct {
	db *sql.DB
}

func NewAccountStore(db *sql.DB) uc.AccountStore {
	return &accountStore{
		db: db,
	}
}

func (as *accountStore) Create(a d.Account) (*d.Account, error) {
	sqlStatement := `INSERT INTO accounts (
		id, 
		user_id, 
		name, 
		description, 
		type, 
		balance_value, 
		balance_currency
	) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	if _, err := as.db.Exec(sqlStatement, a.ID, a.User.ID, a.Name, a.Description, a.Type, a.BalanceValue, a.BalanceCurrency); err != nil {
		return nil, err
	}
	return as.FindByIDAndUser(a.ID, a.User.ID)
}

func (as *accountStore) FindAll(uID d.UserID) (accounts []d.Account, err error) {
	sqlStatement := `SELECT 
		id,
		name,
		description,
		type,
		balance_value,
		balance_currency
	FROM accounts
	WHERE user_id=$1`
	rows, err := as.db.Query(sqlStatement, uID)
	defer rows.Close()
	if err != nil && err != sql.ErrNoRows {
		err = d.ErrInternalError
		return
	}
	for rows.Next() {
		var a d.Account
		err = rows.Scan(&a.ID, &a.Name, &a.Description, &a.Type, &a.BalanceValue, &a.BalanceCurrency)
		if err != nil {
			err = d.ErrInternalError
			return
		}
		accounts = append(accounts, a)
	}
	return
}

func (as *accountStore) FindByIDAndUser(id d.AccountID, uID d.UserID) (*d.Account, error) {
	var account d.Account
	sqlStatement := `SELECT 
		id, 
		name, 
		description, 
		type, 
		balance_value, 
		balance_currency 
	FROM accounts 
	WHERE id=$1 AND user_id=$2`
	if err := as.db.QueryRow(sqlStatement, id, uID).
		Scan(&account.ID, &account.Name, &account.Description, &account.Type, &account.BalanceValue, &account.BalanceCurrency); err != nil {
		switch err {
		case sql.ErrNoRows:
			err = d.ErrNotFound
		default:
			err = d.ErrInternalError
		}
		return nil, err
	}
	return &account, nil
}

func (as *accountStore) FindByUserAndName(uID d.UserID, name string) (*d.Account, error) {
	var account d.Account
	sqlStatement := `SELECT 
		id, 
		name, 
		description, 
		type, 
		balance_value, 
		balance_currency 
	FROM accounts 
	WHERE user_id=$1 AND name=$2`
	if err := as.db.QueryRow(sqlStatement, uID, name).
		Scan(&account.ID, &account.Name, &account.Description, &account.Type, &account.BalanceValue, &account.BalanceCurrency); err != nil {
		switch err {
		case sql.ErrNoRows:
			err = d.ErrNotFound
		default:
			err = d.ErrInternalError
		}
		return nil, err
	}
	return &account, nil
}
