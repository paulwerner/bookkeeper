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
	UserRegister(cmd UserRegisterCmd) (user *d.User, token string, err error)
	UserLogin(cmd UserLoginCmd) (user *d.User, token string, err error)
	UserFind(id d.UserID) (user *d.User, token string, err error)
}

type AppLogic interface {
	AppConfig() (config *d.AppConfig, err error)
}

type AccountLogic interface {
	AccountCreate(cmd AccountCreateCmd) (account *d.Account, err error)
	AccountFind(id d.AccountID, uID d.UserID) (account *d.Account, err error)
}

type TransactionLogic interface {
	TransactionCreate(cmd TransactionCreateCmd) (transaction *d.Transaction, err error)
	TransactionFind(id d.TransactionID, aID d.AccountID) (transaction *d.Transaction, err error)
}

type HandlerConstructor struct {
	AuthHandler   AuthHandler
	UserRW        UserRW
	AccountRW     AccountRW
	TransactionRW TransactionRW
}

func (self HandlerConstructor) New() Handler {
	if self.AuthHandler == nil {
		log.Fatal("missing AuthHandler")
	}
	if self.UserRW == nil {
		log.Fatal("missing UserRW")
	}
	if self.AccountRW == nil {
		log.Fatal("missing AccountRW")
	}
	if self.TransactionRW == nil {
		log.Fatal("missing TransactionRW")
	}
	return interactor{
		authHandler:   self.AuthHandler,
		userRW:        self.UserRW,
		accountRW:     self.AccountRW,
		transactionRW: self.TransactionRW,
	}
}
