package handlers

import (
	"time"

	"github.com/aixoio/speed-notes/server/tools"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type login_request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var dat login_request

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

	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	user, err := tools.UserFromUsername(dat.Username, DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "PostgreSQL cannot find user"})
	}

	if !tools.VerifyPassword(user.Password_hash, dat.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Bad Password"})
	}

	claims := jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add((time.Hour * 24) * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwt_token, err := token.SignedString(CFG.Jwt_secret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot not create JWT"})
	}

	return c.JSON(fiber.Map{"jwt": jwt_token})
}
