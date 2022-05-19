package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// NewJWT creates a new JSON Web Token middleware
func NewJWT() fiber.Handler {
	return jwtware.New(
		jwtware.Config{
			SigningKey: []byte(os.Getenv("JWT_SECRET")),
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				if err.Error() == "Missing or malformed JWT" {
					return c.Status(fiber.StatusForbidden).
						JSON(fiber.Map{
							"errors": fiber.Map{
								"msg": "missing or malformed JWT",
							},
						})

				} else {
					return c.Status(fiber.StatusUnauthorized).
						JSON(fiber.Map{
							"errors": fiber.Map{
								"msg": "invalid or expired JWT",
							},
						})
				}
			},
		},
	)
}
