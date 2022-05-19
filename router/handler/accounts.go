package handler

import (
	"github.com/gofiber/fiber/v2"
	d "github.com/paulwerner/bookkeeper/domain"
	"github.com/paulwerner/bookkeeper/utils"
)

func (h *Handler) AccountsCreate(c *fiber.Ctx) error {
	token, err := getJWT(c.Get("Authorization"))
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}

	uID, err := utils.GetUserIDFromToken(token)
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}

	u, _, err := h.useCases.UserFind(uID)
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}

	var acr accountCreateRequest
	var a d.Account
	if err := acr.bind(c, *u, &a); err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}

	newAccount, err := h.useCases.AccountCreate(a)
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}
	return c.Status(fiber.StatusCreated).
		JSON(newAccountCreateResponse(newAccount))
}

func (h *Handler) AccountsGet(c *fiber.Ctx) error {
	token, err := getJWT(c.Get("Authorization"))
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}

	uID, err := utils.GetUserIDFromToken(token)
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}
	accounts, err := h.useCases.AccountsFind(uID)
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}
	return c.Status(fiber.StatusOK).
		JSON(newAccountsGetResponse(accounts))
}
