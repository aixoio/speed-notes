package handlers

import (
	"strconv"

	"github.com/aixoio/speed-notes/server/tools"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func DeleteNoteHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	uid := uint32(claims["id"].(float64))

	note_id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request id"})
	}

	note, err := tools.NoteByID(uint32(note_id), DB)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "PostgreSQL cannot find note"})
	}

	if note.User_id != uid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "You cannot access this note"})
	}

	err = tools.DeleteNoteByID(note.Id, DB)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "PostgreSQL cannot delete note"})
	}

	return c.JSON(fiber.Map{
		"status": "DELETED",
	})
}
