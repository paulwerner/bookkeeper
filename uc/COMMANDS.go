package uc

import d "github.com/paulwerner/bookkeeper/domain"

type UserLoginCmd struct {
	name     string
	password string
}

type UserRegisterCmd struct {
	id d.UserID
	UserLoginCmd
}

type AccountCreateCmd struct {
	id                     d.AccountID
	uID                    d.UserID
	name                   string
	description            *string
	accountType            d.AccountType
	currentBalanceValue    int64
	currentBalanceCurrency string
}

type TransactionCreateCmd struct {
	id             d.TransactionID
	aID      d.AccountID
	description    *string
	amountValue    int64
	amountCurrency string
}
