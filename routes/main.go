package router

import "github.com/gofiber/fiber/v2"

// Registers func
func Registers(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 2")
	})

	app.Static("/", "../public")
}
