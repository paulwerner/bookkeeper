package utils

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	d "github.com/paulwerner/bookkeeper/domain"
)

func ClearDB(db *sql.DB) {
	db.Exec("DELETE FROM transactions")
	db.Exec("DELETE FROM accounts")
	db.Exec("DELETE FROM users")
}

func RandomUserID() d.UserID {
	return d.UserID(uuid.New().String())
}

func RandomAccountID() d.AccountID {
	return d.AccountID(uuid.New().String())
}

func RandomTransactionID() d.TransactionID {
	return d.TransactionID(uuid.New().String())
}

func PopulateUser(u *d.User, db *sql.DB) {
	sqlStatement := `INSERT INTO users (id, name, password) VALUES ($1, $2, $3)`
	_, err := db.Exec(sqlStatement, u.ID, u.Name, u.Password)
	if err != nil {
		log.Fatal(err)
	}
}

func PopulateAccount(a *d.Account, db *sql.DB) {
	sqlStatement := `INSERT INTO accounts (
		id, 
		user_id, 
		name, 
		description, 
		type, 
		balance_value, 
		balance_currency
	) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := db.Exec(sqlStatement, a.ID, a.User.ID, a.Name, a.Description, a.Type, a.BalanceValue, a.BalanceCurrency)
	if err != nil {
		log.Fatal(err)
	}
}
