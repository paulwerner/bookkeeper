package store

import (
	"database/sql"

	d "github.com/paulwerner/bookkeeper/domain"
	"github.com/paulwerner/bookkeeper/uc"
)

type userStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) uc.UserRW {
	return &userStore{
		db: db,
	}
}

func (self *userStore) Create(id d.UserID, name, password string) (user *d.User, err error) {
	sqlStatement := `INSERT INTO users (id, name, password) VALUES ($1, $2, $3)`
	if _, err = self.db.Exec(sqlStatement, id, name, password); err != nil {
		err = d.ErrInternalError
		return
	}
	user, err = self.FindByID(id)
	return
}

func (self *userStore) FindByID(id d.UserID) (user *d.User, err error) {
	sqlStatement := `SELECT id, name, password FROM users WHERE id=$1`
	if err = self.db.QueryRow(sqlStatement, id).Scan(&user); err != nil {
		switch err {
		case sql.ErrNoRows:
			err = d.ErrNotFound
		default:
			err = d.ErrInternalError
		}
	}
	return
}

func (self *userStore) FindByName(name string) (user *d.User, err error) {
	sqlStatement := `SELECT id, name, password FROM users WHERE name=$1`
	if err = self.db.QueryRow(sqlStatement, name).Scan(&user); err != nil {
		switch err {
		case sql.ErrNoRows:
			err = d.ErrNotFound
		default:
			err = d.ErrInternalError
		}
	}
	return
}
