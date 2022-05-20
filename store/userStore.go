package store

import (
	"database/sql"

	d "github.com/paulwerner/bookkeeper/pkg/domain"
	"github.com/paulwerner/bookkeeper/pkg/uc"
)

type userStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) uc.UserStore {
	return &userStore{
		db: db,
	}
}

func (us *userStore) Create(id d.UserID, name, password string) (*d.User, error) {
	sqlStatement := `INSERT INTO users (id, name, password) VALUES ($1, $2, $3)`
	if _, err := us.db.Exec(sqlStatement, id, name, password); err != nil {
		err = d.ErrInternalError
		return nil, err
	}
	return us.FindByID(id)
}

func (us *userStore) Exists(id d.UserID) (bool, error) {
	sqlStatement := `SELECT 1 FROM users WHERE id=$1`
	if err := us.db.QueryRow(sqlStatement, id).Err(); err != nil {
		switch err {
		case sql.ErrNoRows:
			err = nil
		default:
			err = d.ErrInternalError
		}
		return false, err
	}
	return true, nil
}

func (us *userStore) FindByID(id d.UserID) (*d.User, error) {
	var user d.User
	sqlStatement := `SELECT id, name, password FROM users WHERE id=$1`
	if err := us.db.QueryRow(sqlStatement, id).Scan(&user.ID, &user.Name, &user.Password); err != nil {
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

func (us *userStore) FindByName(name string) (*d.User, error) {
	var user d.User
	sqlStatement := `SELECT id, name, password FROM users WHERE name=$1`
	if err := us.db.QueryRow(sqlStatement, name).Scan(&user.ID, &user.Name, &user.Password); err != nil {
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
