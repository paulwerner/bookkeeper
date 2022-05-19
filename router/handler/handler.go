package handler

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	d "github.com/paulwerner/bookkeeper/domain"
	"github.com/paulwerner/bookkeeper/uc"
	"github.com/paulwerner/bookkeeper/utils"
)

type Handler struct {
	useCases uc.Handler
}

func NewHandler(
	aH uc.AuthHandler,
	uRW uc.UserRW,
	aRW uc.AccountRW,
	txRW uc.TransactionRW,
) *Handler {
	return &Handler{
		useCases: uc.HandlerConstructor{
			AuthHandler:   aH,
			UserRW:        uRW,
			AccountRW:     aRW,
			TransactionRW: txRW,
		}.New(),
	}
}

func getJWT(authHeader string) (string, error) {
	splitted := strings.Split(authHeader, "Bearer ")
	if len(splitted) != 2 {
		return "", d.ErrInvalidAccessToken
	}
	return splitted[1], nil
}

func (h *Handler) getUserIDFromRequest(c *fiber.Ctx) (d.UserID, error) {
	token, err := getJWT(c.Get("Authorization"))
	if err != nil {
		return "", err
	}
	return utils.GetUserIDFromToken(token)
}
