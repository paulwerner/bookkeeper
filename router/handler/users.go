package handler

import (
	"github.com/gofiber/fiber/v2"
	d "github.com/paulwerner/bookkeeper/domain"
	"github.com/paulwerner/bookkeeper/utils"
)

func (h *Handler) SignUp(c *fiber.Ctx) error {
	var u d.User
	req := userSignUpRequest{}
	if err := req.bind(c, &u); err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}
	user, token, err := h.useCases.UserRegister(
		utils.RandomUserID(),
		u.Name,
		u.Password,
	)
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}
	return c.Status(fiber.StatusCreated).
		JSON(newUserSignUpResponse(user, token))
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var u d.User
	req := &userLoginRequest{}
	if err := req.bind(c, &u); err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}
	user, token, err := h.useCases.UserLogin(u.Name, u.Password)
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}
	return c.Status(fiber.StatusOK).
		JSON(newUserLoginResponse(user, token))
}
