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
) (*d.Account, error) {
	sqlStatement := `INSERT INTO accounts (
		id, 
		user_id, 
		name, 
		description, 
		type, 
		balance_value, 
		balance_currency
	) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	if _, err := self.db.Exec(sqlStatement, id, uID, name, description, accountType, currentBalanceValue, currentBalanceCurrency); err != nil {
		return nil, err
	}
	return self.FindByIDAndUser(id, uID)
}

func (self *accountStore) FindByIDAndUser(id d.AccountID, uID d.UserID) (*d.Account, error) {
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
	if err := self.db.QueryRow(sqlStatement, id, uID).
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

func (self *accountStore) FindByUserAndName(uID d.UserID, name string) (*d.Account, error) {
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
	if err := self.db.QueryRow(sqlStatement, uID, name).
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
