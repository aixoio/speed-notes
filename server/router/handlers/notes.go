package handlers

import (
	"github.com/aixoio/speed-notes/server/tools"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func NotesHandlers(c *fiber.Ctx) error {
	c.Accepts("application/json")

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	uid := uint32(claims["id"].(float64))

	notes, err := tools.AllNotesByUserID(uid, DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "PostgreSQL cannot find notes"})
	}

	return c.JSON(notes)
}
