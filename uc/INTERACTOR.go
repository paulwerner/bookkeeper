package uc

import d "github.com/paulwerner/bookkeeper/domain"

type interactor struct {
	authHandler   AuthHandler
	userRW        UserRW
	accountRW     AccountRW
	transactionRW TransactionRW
}

type AuthHandler interface {
	GenUserToken(id d.UserID) (token string, err error)
	GetUserID(token string) (id d.UserID, err error)
	EncryptPassword(password string) (encryptedPassword string, err error)
	CheckPassword(hashedPassword, password string) (ok bool)
}

type UserRW interface {
	Create(id d.UserID, name, password string) (user *d.User, err error)
	FindByID(id d.UserID) (user *d.User, err error)
	FindByName(name string) (user *d.User, err error)
}

type AccountRW interface {
	Create(
		id d.AccountID,
		uID d.UserID,
		name string,
		description *string,
		accountType d.AccountType,
		currentBalanceValue int64,
		currentBalanceCurrency string,
	) (account *d.Account, err error)
	FindAll(uID d.UserID) (accounts []d.Account, err error)
	FindByUserAndName(uID d.UserID, name string) (account *d.Account, err error)
	FindByIDAndUser(id d.AccountID, uID d.UserID) (account *d.Account, err error)
}

type TransactionRW interface {
	Create(
		id d.TransactionID,
		aID d.AccountID,
		description *string,
		amountValue int64,
		amountCurrency string,
	) (tx *d.Transaction, err error)
	FindByIDAndAccount(id d.TransactionID, aID d.AccountID) (tx *d.Transaction, err error)
}
