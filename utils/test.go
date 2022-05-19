package utils

import (
	"database/sql"
	"log"

	d "github.com/paulwerner/bookkeeper/pkg/domain"
)

func ClearDB(db *sql.DB) {
	db.Exec("DELETE FROM transactions")
	db.Exec("DELETE FROM accounts")
	db.Exec("DELETE FROM users")
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

func PopulateTransaction(tx *d.Transaction, db *sql.DB) {
	sqlStatement := `INSERT INTO transactions (
		id, 
		account_id, 
		description, 
		amount, 
		currency
	) VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(sqlStatement, tx.ID, tx.Account.ID, tx.Description, tx.Amount, tx.Currency)
	if err != nil {
		log.Fatal(err)
	}
}
