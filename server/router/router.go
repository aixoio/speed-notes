package router

import (
	"database/sql"
	"fmt"

	"github.com/aixoio/speed-notes/server/env"
	"github.com/aixoio/speed-notes/server/router/handlers"
	"github.com/gofiber/fiber/v2"
)

var DB *sql.DB

func StartRouter(cfg env.Config, db *sql.DB) {
	DB = db

	app := fiber.New()

	app.Post("/signup", handlers.SignupHandler)

	app.Listen(fmt.Sprintf(":%s", cfg.Port))
}
