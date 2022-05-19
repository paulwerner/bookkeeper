package handler

import "github.com/paulwerner/bookkeeper/uc"

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
