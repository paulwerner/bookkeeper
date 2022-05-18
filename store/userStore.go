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

func (self *userStore) Create(id d.UserID, name, password string) (*d.User, error) {
	sqlStatement := `INSERT INTO users (id, name, password) VALUES ($1, $2, $3)`
	if _, err := self.db.Exec(sqlStatement, id, name, password); err != nil {
		err = d.ErrInternalError
		return nil, err
	}
	return self.FindByID(id)
}

func (self *userStore) FindByID(id d.UserID) (*d.User, error) {
	var user d.User
	sqlStatement := `SELECT id, name, password FROM users WHERE id=$1`
	if err := self.db.QueryRow(sqlStatement, id).Scan(&user.ID, &user.Name, &user.Password); err != nil {
		switch err {
		case sql.ErrNoRows:
			err = d.ErrNotFound
		default:
			err = d.ErrInternalError
		}
		return nil, err
	}
	return &user, nil
}

func (self *userStore) FindByName(name string) (*d.User, error) {
	var user d.User
	sqlStatement := `SELECT id, name, password FROM users WHERE name=$1`
	if err := self.db.QueryRow(sqlStatement, name).Scan(&user.ID, &user.Name, &user.Password); err != nil {
		switch err {
		case sql.ErrNoRows:
			err = d.ErrNotFound
		default:
			err = d.ErrInternalError
		}
		return nil, err
	}
	return &user, nil
}
