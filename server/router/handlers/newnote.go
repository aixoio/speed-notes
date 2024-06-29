package handlers

import (
	"github.com/aixoio/speed-notes/server/tools"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type new_note_request struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewNoteHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	uid := uint32(claims["id"].(float64))

	var dat new_note_request

	if err := c.BodyParser(&dat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request body"})
	}

	note_id, err := tools.InsertNewNote(uid, dat.Title, dat.Content, DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "PostgreSQL cannot insert new note"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"note_id": note_id})
}
