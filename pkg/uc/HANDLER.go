package uc

import (
	"log"

	d "github.com/paulwerner/bookkeeper/pkg/domain"
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
	TransactionCreate(tx d.Transaction) (transaction *d.Transaction, err error)
	TransactionFind(id d.TransactionID, aID d.AccountID) (transaction *d.Transaction, err error)
	TransactionsFind(aID d.AccountID) (transactions []d.Transaction, err error)
}

type HandlerConstructor struct {
	AuthHandler      AuthHandler
	UserStore        UserStore
	AccountStore     AccountStore
	TransactionStore TransactionStore
}

func (hC HandlerConstructor) New() Handler {
	if hC.AuthHandler == nil {
		log.Fatal("missing AuthHandler")
	}
	if hC.UserStore == nil {
		log.Fatal("missing UserStore")
	}
	if hC.AccountStore == nil {
		log.Fatal("missing AccountStore")
	}
	if hC.TransactionStore == nil {
		log.Fatal("missing TransactionStore")
	}
	return interactor{
		authHandler:      hC.AuthHandler,
		userStore:        hC.UserStore,
		accountStore:     hC.AccountStore,
		transactionStore: hC.TransactionStore,
	}
}
