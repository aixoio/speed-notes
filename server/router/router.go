package router

import (
	"database/sql"
	"fmt"

	"github.com/aixoio/speed-notes/server/env"
	"github.com/aixoio/speed-notes/server/router/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func StartRouter(cfg env.Config, db *sql.DB) {
	handlers.DB = db

	app := fiber.New()

	app.Use(logger.New())

	app.Post("/signup", handlers.SignupHandler)
	app.Post("/login", handlers.LoginHandler)

	app.Listen(fmt.Sprintf(":%s", cfg.Port))
}
