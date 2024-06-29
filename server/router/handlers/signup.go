package handlers

import (
	"github.com/aixoio/speed-notes/server/data"
	"github.com/aixoio/speed-notes/server/tools"
	"github.com/gofiber/fiber/v2"
)

type signup_request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignupHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var dat signup_request

	if err := c.BodyParser(&dat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request body"})
	}

	if !tools.IsUsernameValid(dat.Username) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "The username is invalid"})
	}

	exists, err := tools.UserExists(dat.Username, DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot check username"})
	}

	if exists {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "The username is taken"})
	}

	pwd_hash, err := tools.HashPassword(dat.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot hash password"})
	}

	err = tools.UserInsert(data.User{Username: dat.Username, Password_hash: pwd_hash}, DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "PostgreSQL cannot insert user"})
	}

	return c.Status(fiber.StatusCreated).JSON(map[string]string{
		"status": "CREATED",
	})
}
