package handler

import (
	"github.com/gofiber/fiber/v2"
	d "github.com/paulwerner/bookkeeper/pkg/domain"
	"github.com/paulwerner/bookkeeper/utils"
)

func (h *Handler) SignUp(c *fiber.Ctx) error {
	// bind sign up request body
	var u d.User
	req := userSignUpRequest{}
	if err := req.bind(c, &u); err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}
	// register new user
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
	// bind login request body
	var u d.User
	req := &userLoginRequest{}
	if err := req.bind(c, &u); err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}
	// login user
	user, token, err := h.useCases.UserLogin(u.Name, u.Password)
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}
	return c.Status(fiber.StatusOK).
		JSON(newUserLoginResponse(user, token))
}
