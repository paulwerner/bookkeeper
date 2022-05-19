package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) ConfigGet(c *fiber.Ctx) error {
	conf, err := h.useCases.AppConfig()
	if err != nil {
		errBody, sc := newErrorResponse(err)
		return c.Status(sc).JSON(errBody)
	}
	return c.Status(fiber.StatusOK).JSON(newAppConfigResponse(conf))
}
