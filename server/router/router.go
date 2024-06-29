package router

import (
	"github.com/aixoio/speed-notes/server/env"
	"github.com/gofiber/fiber/v2"
)

func StartRouter(cfg env.Config) {
	app := fiber.New()

	app.Listen(":8080")
}
