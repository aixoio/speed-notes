package router

import "github.com/gofiber/fiber/v2"

func StartRouter() {
	app := fiber.New()

	app.Listen(":8080")
}
