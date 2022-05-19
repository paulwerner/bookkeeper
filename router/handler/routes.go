package handler

import (
	"github.com/gofiber/fiber/v2"
	mw "github.com/paulwerner/bookkeeper/router/middleware"
)

func (h *Handler) Register(r *fiber.App) {
	jwtMW := mw.NewJWT()
	v1 := r.Group("/api")
	guestUsers := v1.Group("/users")
	guestUsers.Post("", h.SignUp)
	guestUsers.Post("/login", h.Login)

	app := v1.Group("/app", jwtMW)
	app.Get("/config", h.AppConfig)

	accounts := v1.Group("/accounts", jwtMW)
	accounts.Post("", h.AccountsCreate)
	accounts.Get("", h.AccountsGet)
	accounts.Get("/:account_id", h.AccountGet)
}
