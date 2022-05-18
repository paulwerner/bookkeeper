package utils

import (
	"database/sql"
	"log"

	d "github.com/paulwerner/bookkeeper/domain"
)

func ClearDB(db *sql.DB) {
	db.Exec("DELETE FROM users")
	db.Exec("DELETE FROM accounts")
	db.Exec("DELETE FROM transactions")
}

func PopulateUser(u *d.User, db *sql.DB) {
	sqlStatement := `INSERT INTO users (id, name, password) VALUES ($1, $2, $3)`
	_, err := db.Exec(sqlStatement, u.ID, u.Name, u.Password)
	if err != nil {
		log.Fatal(err)
	}
}
