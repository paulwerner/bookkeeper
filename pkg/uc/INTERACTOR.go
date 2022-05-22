package uc

import d "github.com/paulwerner/bookkeeper/pkg/domain"

type interactor struct {
	authHandler      AuthHandler
	userStore        UserStore
	accountStore     AccountStore
	transactionStore TransactionStore
}

type AuthHandler interface {
	GenUserToken(id d.UserID) (token string, err error)
	GetUserID(token string) (id d.UserID, err error)
	EncryptPassword(password string) (encryptedPassword string, err error)
	CheckPassword(hashedPassword, password string) (ok bool)
}

type UserStore interface {
	Create(id d.UserID, name, password string) (user *d.User, err error)
	Exists(id d.UserID) (ok bool, err error)
	FindByID(id d.UserID) (user *d.User, err error)
	FindByName(name string) (user *d.User, err error)
}

type AccountStore interface {
	Create(a d.Account) (account *d.Account, err error)
	Update(a d.Account) (account *d.Account, err error)
	FindAll(uID d.UserID) (accounts []d.Account, err error)
	FindByUserAndName(uID d.UserID, name string) (account *d.Account, err error)
	FindByIDAndUser(id d.AccountID, uID d.UserID) (account *d.Account, err error)
}

type TransactionStore interface {
	Create(tx d.Transaction) (transaction *d.Transaction, err error)
	FindByIDAndAccount(id d.TransactionID, aID d.AccountID) (transaction *d.Transaction, err error)
	FindAll(aID d.AccountID) (txs []d.Transaction, err error)
}
