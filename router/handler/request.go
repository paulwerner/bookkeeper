package handler

import (
	"github.com/gofiber/fiber/v2"
	d "github.com/paulwerner/bookkeeper/domain"
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