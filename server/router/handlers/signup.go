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
		return fiber.NewError(fiber.StatusBadRequest, "Cannot parse request body")
	}

	if !tools.IsUsernameValid(dat.Username) {
		return fiber.NewError(fiber.StatusBadRequest, "The username is invalid")
	}

	exists, err := tools.UserExists(dat.Username, DB)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Cannot check username")
	}

	if exists {
		return fiber.NewError(fiber.StatusConflict, "The username is taken")
	}

	pwd_hash, err := tools.HashPassword(dat.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Cannot hash password")
	}

	err = tools.UserInsert(data.User{Username: dat.Username, Password_hash: pwd_hash}, DB)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "PostgreSQL cannot insert user")
	}

	return c.Status(fiber.StatusCreated).JSON(map[string]string{
		"success": "CREATED",
	})
}
