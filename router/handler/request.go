package handler

import (
	"github.com/gofiber/fiber/v2"
	d "github.com/paulwerner/bookkeeper/domain"
	"github.com/paulwerner/bookkeeper/utils"
)

type userSignUpRequest struct {
	User struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	} `json:"user"`
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
	a = d.NewAccount(
		utils.RandomAccountID(),
		u,
		r.Account.Name,
		r.Account.Description,
		r.Account.Type,
		r.Account.CurrentBalance.Value,
		r.Account.CurrentBalance.Currency,
	)
	return nil
}
