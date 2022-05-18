package store

import (
	"database/sql"

	d "github.com/paulwerner/bookkeeper/domain"
	"github.com/paulwerner/bookkeeper/uc"
)

type accountStore struct {
	db *sql.DB
}

func NewAccountStore(db *sql.DB) uc.AccountRW {
	return &accountStore{
		db: db,
	}
}

func (self *accountStore) Create(
	id d.AccountID,
	uID d.UserID,
	name string,
	description *string,
	accountType d.AccountType,
	currentBalanceValue int64,
	currentBalanceCurrency string,
) (account *d.Account, err error) {
	sqlStatement := `INSERT INTO accounts (
		id, 
		user_id, 
		name, 
		description, 
		account_type, 
		balance_value, 
		balance_currency
	) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	if _, err = self.db.Exec(sqlStatement, id, uID, name, description, accountType, currentBalanceValue, currentBalanceCurrency); err != nil {
		err = d.ErrInternalError
		return
	}
	account, err = self.FindByIDAndUser(id, uID)
	return
}

func (self *accountStore) FindByIDAndUser(id d.AccountID, uID d.UserID) (account *d.Account, err error) {
	sqlStatement := `SELECT 
		id, 
		name, 
		description, 
		account_type, 
		balance_value, 
		balance_currency 
	FROM accounts 
	WHERE id=$1 AND user_id=$2`
	if err = self.db.QueryRow(sqlStatement, id, uID).Scan(&account); err != nil {
		switch err {
		case sql.ErrNoRows:
			err = d.ErrNotFound
		default:
			err = d.ErrInternalError
		}
	}
	return
}

func (self *accountStore) FindByUserAndName(uID d.UserID, name string) (account *d.Account, err error) {
	sqlStatement := `SELECT 
		id, 
		name, 
		description, 
		account_type, 
		balance_value, 
		balance_currency 
	FROM accounts 
	WHERE user_id=$1 AND name=$2`
	if err = self.db.QueryRow(sqlStatement, uID, name).Scan(&account); err != nil {
		switch err {
		case sql.ErrNoRows:
			err = d.ErrNotFound
		default:
			err = d.ErrInternalError
		}
	}
	return
}
