package handler

import (
	"github.com/gofiber/fiber/v2"
	d "github.com/paulwerner/bookkeeper/pkg/domain"
)

func (h *Handler) TransactionGet(c *fiber.Ctx) error {
	// get user id from request
	uID, err := h.getUserIDFromRequest(c)
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}
	// check if user exists
	ok := h.useCases.UserExists(uID)
	if !ok {
		errBody, sc := newErrorResponse(d.ErrNotFound)
		return c.Status(sc).JSON(errBody)
	}
	// find transaction for given id and account
	aID := d.AccountID(c.Params("account_id"))
	txID := d.TransactionID(c.Params("transaction_id"))
	tx, err := h.useCases.TransactionFind(txID, aID)
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}
	return c.Status(fiber.StatusOK).
		JSON(newTransactionGetResponse(tx))
}

func (h *Handler) TransactionsGet(c *fiber.Ctx) error {
	// get user id from request
	uID, err := h.getUserIDFromRequest(c)
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}
	// check if user exists
	ok := h.useCases.UserExists(uID)
	if !ok {
		errBody, sc := newErrorResponse(d.ErrNotFound)
		return c.Status(sc).JSON(errBody)
	}
	// find transaction
	aID := d.AccountID(c.Params("account_id"))
	txs, err := h.useCases.TransactionsFind(aID)
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}
	return c.Status(fiber.StatusOK).
		JSON(newTransactionsGetResponse(txs))
}

func (h *Handler) TransactionCreate(c *fiber.Ctx) error {
	// get user id from request
	uID, err := h.getUserIDFromRequest(c)
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}
	// check if user exists
	ok := h.useCases.UserExists(uID)
	if !ok {
		errBody, sc := newErrorResponse(d.ErrNotFound)
		return c.Status(sc).JSON(errBody)
	}
	// get account for given id
	aID := d.AccountID(c.Params("account_id"))
	a, err := h.useCases.AccountFind(aID, uID)
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}
	// bind create transaction request
	var ctr transactionCreateRequest
	var tx d.Transaction
	if err := ctr.bind(c, *a, &tx); err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}
	// create new transaction for account
	newTx, err := h.useCases.TransactionCreate(tx)
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}
	return c.Status(fiber.StatusCreated).
		JSON(newTransactionGetResponse(newTx))
}
