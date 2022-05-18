package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	d "github.com/paulwerner/bookkeeper/domain"
)

var tokenTimeToLive = 2 * time.Hour

// GenerateJWT generates a new JWT token for a given user ID
func GenerateJWT(id d.UserID) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(tokenTimeToLive).Unix()

	secret := []byte("top-secret")
	t, err := token.SignedString(secret)
	if err != nil {
		panic(err)
	}
	return t
}

// UserIDFromToken returns the user id from the token
func UserIDFromToken(c *fiber.Ctx) (*d.UserID, error) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	uid := d.UserID(id)
	return &uid, nil
}
