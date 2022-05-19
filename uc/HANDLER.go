package uc

import (
	"log"

	d "github.com/paulwerner/bookkeeper/domain"
)

type Handler interface {
	UserLogic
	AppLogic
	AccountLogic
	TransactionLogic
}

type UserLogic interface {
	UserRegister(id d.UserID, name, password string) (user *d.User, token string, err error)
	UserLogin(name, password string) (user *d.User, token string, err error)
	UserExists(id d.UserID) (ok bool)
	UserFind(id d.UserID) (user *d.User, token string, err error)
}

type AppLogic interface {
	AppConfig() (config *d.AppConfig, err error)
}

type AccountLogic interface {
	AccountCreate(a d.Account) (account *d.Account, err error)
	AccountsFind(uID d.UserID) (account []d.Account, err error)
	AccountFind(id d.AccountID, uID d.UserID) (account *d.Account, err error)
}

type TransactionLogic interface {
	TransactionCreate(
		id d.TransactionID,
		aID d.AccountID,
		description *string,
		amount int64,
		currency string,
	) (transaction *d.Transaction, err error)
	TransactionFind(id d.TransactionID, aID d.AccountID) (transaction *d.Transaction, err error)
}

type HandlerConstructor struct {
	AuthHandler   AuthHandler
	UserRW        UserRW
	AccountRW     AccountRW
	TransactionRW TransactionRW
}

func (hC HandlerConstructor) New() Handler {
	if hC.AuthHandler == nil {
		log.Fatal("missing AuthHandler")
	}
	if hC.UserRW == nil {
		log.Fatal("missing UserRW")
	}
	if hC.AccountRW == nil {
		log.Fatal("missing AccountRW")
	}
	if hC.TransactionRW == nil {
		log.Fatal("missing TransactionRW")
	}
	return interactor{
		authHandler:   hC.AuthHandler,
		userRW:        hC.UserRW,
		accountRW:     hC.AccountRW,
		transactionRW: hC.TransactionRW,
	}
}
