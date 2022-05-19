package handler

import (
	"github.com/Rhymond/go-money"
	d "github.com/paulwerner/bookkeeper/domain"
)

// User
type userSignUpResponse struct {
	User struct {
		Name string `json:"name"`
	} `json:"user"`
	Token string `json:"token"`
}

func newUserSignUpResponse(u *d.User, token string) *userSignUpResponse {
	var resp userSignUpResponse
	resp.User.Name = u.Name
	resp.Token = token
	return &resp
}

type userLoginResponse struct {
	User struct {
		Name string `json:"name"`
	} `json:"user"`
	Token string `json:"access_token"`
}

func newUserLoginResponse(u *d.User, token string) *userLoginResponse {
	var resp userLoginResponse
	resp.User.Name = u.Name
	resp.Token = token
	return &resp
}

// App Config
type appConfigResponse struct {
	SupportedAccountTypes []d.AccountType `json:"supported_account_types"`
}

func newAppConfigResponse(c *d.AppConfig) *appConfigResponse {
	return &appConfigResponse{
		SupportedAccountTypes: c.SupportedAccountTypes,
	}
}

// Accounts
type accountResponse struct {
	ID               d.AccountID   `json:"id"`
	Name             string        `json:"name"`
	Description      *string       `json:"description"`
	Type             d.AccountType `json:"type"`
	BalanceFormatted string        `json:"balance_formatted"`
}

func newAccountResponse(a *d.Account) *accountResponse {
	return &accountResponse{
		ID:               a.ID,
		Name:             a.Name,
		Description:      a.Description,
		Type:             a.Type,
		BalanceFormatted: money.New(a.BalanceValue, a.BalanceCurrency).Display(),
	}
}

type accountCreateResponse struct {
	Account accountResponse `json:"account"`
}

func newAccountCreateResponse(a *d.Account) *accountCreateResponse {
	return &accountCreateResponse{*newAccountResponse(a)}
}

type accountGetResponse struct {
	Account accountResponse `json:"account"`
}

func newAccountGetResponse(a *d.Account) *accountGetResponse {
	return &accountGetResponse{*newAccountResponse(a)}
}

type accountsGetResponse struct {
	Accounts []accountResponse `json:"accounts"`
}

func newAccountsGetResponse(accounts []d.Account) *accountsGetResponse {
	var accountsGetResponse accountsGetResponse
	for _, a := range accounts {
		accountsGetResponse.Accounts = append(accountsGetResponse.Accounts, *newAccountResponse(&a))
	}
	return &accountsGetResponse
}

// Transactions
type transactionResponse struct {
	ID              d.TransactionID `json:"id"`
	Description     *string         `json:"description"`
	AmountFormatted string          `json:"amount_formatted"`
}

func newTransactionResponse(tx *d.Transaction) *transactionResponse {
	return &transactionResponse{
		ID:              tx.ID,
		Description:     tx.Description,
		AmountFormatted: money.New(tx.Amount, tx.Currency).Display(),
	}
}

type transactionGetResponse struct {
	Transaction transactionResponse `json:"transaction"`
}

func newTransactionGetResponse(tx *d.Transaction) *transactionGetResponse {
	return &transactionGetResponse{*newTransactionResponse(tx)}
}

type transactionsGetResponse struct {
	Transactions []transactionResponse `json:"transactions"`
}

func newTransactionsGetResponse(txs []d.Transaction) *transactionsGetResponse {
	var tgr transactionsGetResponse
	for _, tx := range txs {
		tgr.Transactions = append(tgr.Transactions, *newTransactionResponse(&tx))
	}
	return &tgr
}
