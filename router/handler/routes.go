package handler

import (
	"github.com/gofiber/fiber/v2"
	mw "github.com/paulwerner/bookkeeper/router/middleware"
)

func (h *Handler) Register(r *fiber.App) {
	jwtMW := mw.NewJWT()
	v1 := r.Group("/api")

	// users
	guestUsers := v1.Group("/users")
	guestUsers.Post("", h.SignUp)
	guestUsers.Post("/login", h.Login)

	// config
	app := v1.Group("/config", jwtMW)
	app.Get("", h.ConfigGet)

	// accounts
	accounts := v1.Group("/accounts", jwtMW)
	accounts.Post("", h.AccountsCreate)
	accounts.Get("", h.AccountsGet)

	account := accounts.Group("/:account_id", jwtMW)
	account.Get("/", h.AccountGet)

	// transactions
	transactions := account.Group("/transactions", jwtMW)
	transactions.Get("", h.TransactionsGet)
	transactions.Post("", h.TransactionCreate)
	transactions.Get("/:transaction_id", h.TransactionGet)

}
