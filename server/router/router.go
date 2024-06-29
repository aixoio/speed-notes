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

	app.Post("/api/signup", handlers.SignupHandler)
	app.Post("/api/login", handlers.LoginHandler)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: cfg.Jwt_secret},
	}))

	app.Get("/api/notes", handlers.NotesHandlers)
	app.Post("/api/notes/new", handlers.NewNoteHandler)
	app.Get("/api/notes/get/:id", handlers.GetNoteHandler)
	app.Post("/api/notes/update/:id", handlers.UpdateNoteHandler)
	app.Post("/api/notes/delete/:id", handlers.DeleteNoteHandler)

	app.Listen(fmt.Sprintf(":%s", cfg.Port))
}
