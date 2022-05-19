package handler

import (
	"github.com/gofiber/fiber/v2"
	d "github.com/paulwerner/bookkeeper/domain"
	"github.com/paulwerner/bookkeeper/utils"
)

func (h *Handler) TransactionGet(c *fiber.Ctx) error {
	// get auth token from header
	token, err := getJWT(c.Get("Authorization"))
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}

	// check if user exists
	uID, err := utils.GetUserIDFromToken(token)
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}
	ok := h.useCases.UserExists(uID)
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}
	if !ok {
		errBody, sc := newErrorResponse(d.ErrNotFound)
		return c.Status(sc).JSON(errBody)
	}

	// find transaction
	aID := d.AccountID(c.Params("account_id"))
	txID := d.TransactionID(c.Params("transaction_id"))
	tx, err := h.useCases.TransactionFind(txID, aID)
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}
	return c.Status(fiber.StatusOK).JSON(newTransactionGetResponse(tx))

}
