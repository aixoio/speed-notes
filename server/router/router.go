package router

import (
	"fmt"

	"github.com/aixoio/speed-notes/server/env"
	"github.com/aixoio/speed-notes/server/router/handlers"
	"github.com/gofiber/fiber/v2"
)

func StartRouter(cfg env.Config) {
	app := fiber.New()

	app.Post("/signup", handlers.SignupHandler)

	app.Listen(fmt.Sprintf(":%s", cfg.Port))
}
