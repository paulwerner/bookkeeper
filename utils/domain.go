package utils

import (
	"github.com/google/uuid"
	d "github.com/paulwerner/bookkeeper/pkg/domain"
)

func RandomUserID() d.UserID {
	return d.UserID(uuid.New().String())
}

func RandomAccountID() d.AccountID {
	return d.AccountID(uuid.New().String())
}

func RandomTransactionID() d.TransactionID {
	return d.TransactionID(uuid.New().String())
}
