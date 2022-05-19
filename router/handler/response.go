package handler

import (
	"github.com/Rhymond/go-money"
	d "github.com/paulwerner/bookkeeper/domain"
)

type userSignUpResponse struct {
	User struct {
		Name string `json:"name"`
	} `json:"user"`
	Token string `json:"token"`
}

func newUserSignUpResponse(u *d.User, token string) *userSignUpResponse {
	resp := userSignUpResponse{}
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
	resp := userLoginResponse{}
	resp.User.Name = u.Name
	resp.Token = token
	return &resp
}

type appConfigResponse struct {
	SupportedAccountTypes []d.AccountType `json:"supported_account_types"`
}

func newAppConfigResponse(c *d.AppConfig) *appConfigResponse {
	return &appConfigResponse{
		SupportedAccountTypes: c.SupportedAccountTypes,
	}
}

type accountCreateResponse struct {
	Account struct {
		ID               d.AccountID   `json:"id"`
		Name             string        `json:"name"`
		Description      *string       `json:"description"`
		Type             d.AccountType `json:"type"`
		BalanceFormatted string        `json:"balance_formatted"`
	} `json:"account"`
}

func newAccountCreateResponse(a *d.Account) *accountCreateResponse {
	resp := accountCreateResponse{}
	resp.Account.ID = a.ID
	resp.Account.Name = a.Name
	resp.Account.Description = a.Description
	resp.Account.BalanceFormatted = money.New(a.BalanceValue, a.BalanceCurrency).Display()
	return &resp
}
