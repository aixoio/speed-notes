package handlers

import "github.com/gofiber/fiber/v2"

func SignupHandler(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"Hello": "world",
	})
}
