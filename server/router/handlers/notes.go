package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func NotesHandlers(c *fiber.Ctx) error {
	c.Accepts("application/json")

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	uid := uint32(claims["id"].(float64))

}
