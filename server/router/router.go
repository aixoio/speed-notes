package router

import (
	"database/sql"
	"fmt"

	"github.com/aixoio/speed-notes/server/env"
	"github.com/aixoio/speed-notes/server/router/handlers"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func StartRouter(cfg env.Config, db *sql.DB) {
	handlers.DB = db
	handlers.CFG = cfg

	app := fiber.New()

	app.Use(logger.New())

	app.Post("/signup", handlers.SignupHandler)
	app.Post("/login", handlers.LoginHandler)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: cfg.Jwt_secret},
	}))

	app.Get("/notes", handlers.NotesHandlers)
	app.Get("/notes/new", handlers.NewNoteHandler)

	app.Listen(fmt.Sprintf(":%s", cfg.Port))
}
