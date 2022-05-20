package handler

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	d "github.com/paulwerner/bookkeeper/pkg/domain"
	"github.com/paulwerner/bookkeeper/pkg/uc"
	"github.com/paulwerner/bookkeeper/utils"
)

type Handler struct {
	useCases uc.Handler
}

func NewHandler(
	aH uc.AuthHandler,
	us uc.UserStore,
	as uc.AccountStore,
	ts uc.TransactionStore,
) *Handler {
	return &Handler{
		useCases: uc.HandlerConstructor{
			AuthHandler:      aH,
			UserStore:        us,
			AccountStore:     as,
			TransactionStore: ts,
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
