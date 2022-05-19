package handler

import (
	"github.com/gofiber/fiber/v2"
	d "github.com/paulwerner/bookkeeper/pkg/domain"
)

func (h *Handler) AccountsCreate(c *fiber.Ctx) error {
	// get user id from request
	uID, err := h.getUserIDFromRequest(c)
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}

	// find user
	u, _, err := h.useCases.UserFind(uID)
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}

	// bind create account request body
	var acr accountCreateRequest
	var a d.Account
	if err := acr.bind(c, *u, &a); err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}

	// create account for user
	newAccount, err := h.useCases.AccountCreate(a)
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}
	return c.Status(fiber.StatusCreated).
		JSON(newAccountCreateResponse(newAccount))
}

func (h *Handler) AccountsGet(c *fiber.Ctx) error {
	// get user id from request
	uID, err := h.getUserIDFromRequest(c)
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}
	// find accounts for user
	accounts, err := h.useCases.AccountsFind(uID)
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}
	return c.Status(fiber.StatusOK).
		JSON(newAccountsGetResponse(accounts))
}

func (h *Handler) AccountGet(c *fiber.Ctx) error {
	// get user id from request
	uID, err := h.getUserIDFromRequest(c)
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}
	// find account for given id and user
	aID := d.AccountID(c.Params("account_id"))
	account, err := h.useCases.AccountFind(aID, uID)
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}
	return c.Status(fiber.StatusOK).
		JSON(newAccountGetResponse(account))
}
