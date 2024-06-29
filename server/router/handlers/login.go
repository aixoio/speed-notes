package handlers

import "github.com/gofiber/fiber/v2"

type login_request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")
}
