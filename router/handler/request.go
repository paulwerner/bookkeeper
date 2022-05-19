package handler

import (
	"github.com/gofiber/fiber/v2"
	d "github.com/paulwerner/bookkeeper/domain"
	"github.com/paulwerner/bookkeeper/utils"
)

// User
type userSignUpRequest struct {
	User userLoginRequest `json:"user"`
}

func (r *userSignUpRequest) bind(c *fiber.Ctx, u *d.User) error {
	if err := c.BodyParser(r); err != nil {
		return d.ErrInvalidEntity
	}
	u.Name = r.User.Name
	u.Password = r.User.Password
	return nil
}

type userLoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (r *userLoginRequest) bind(c *fiber.Ctx, u *d.User) error {
	if err := c.BodyParser(r); err != nil {
		return d.ErrInvalidEntity
	}
	u.Name = r.Name
	u.Password = r.Password
	return nil
}

// Account
type accountCreateRequest struct {
	Account struct {
		Name           string        `json:"name"`
		Description    *string       `json:"description"`
		Type           d.AccountType `json:"type"`
		CurrentBalance struct {
			Value    int64  `json:"value"`
			Currency string `json:"currency"`
		}
	} `json:"account"`
}

func (r *accountCreateRequest) bind(c *fiber.Ctx, u d.User, a *d.Account) error {
	if err := c.BodyParser(r); err != nil {
		return d.ErrInvalidEntity
	}
	a.ID = utils.RandomAccountID()
	a.User = u
	a.Name = r.Account.Name
	a.Description = r.Account.Description
	a.Type = r.Account.Type
	a.BalanceValue = r.Account.CurrentBalance.Value
	a.BalanceCurrency = r.Account.CurrentBalance.Currency
	return nil
}

// Transaction
type transactionCreateRequest struct {
	Transaction struct {
		Description *string `json:"description"`
		Amount      int64   `json:"amount"`
		Currency    string  `json:"currency"`
	} `json:"transaction"`
}

func (r *transactionCreateRequest) bind(c *fiber.Ctx, a d.Account, tx *d.Transaction) error {
	if err := c.BodyParser(r); err != nil {
		return d.ErrInvalidEntity
	}
	tx.ID = utils.RandomTransactionID()
	tx.Account = a
	tx.Description = r.Transaction.Description
	tx.Amount = r.Transaction.Amount
	tx.Currency = r.Transaction.Currency
	return nil
}
