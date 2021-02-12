package router

import (
	handler "github/Khemmathat321/goCRUD/handlers"

	"github.com/gofiber/fiber/v2"
)

// Registers func
func Registers(app *fiber.App) {
	app.Static("/public", "../public")

	// Middleware
	api := app.Group("/api")

	User(api)
}

// User func
func User(router fiber.Router) {
	router = router.Group("/users")

	// Basic resouce routes
	router.Get("/:id?", handler.Get)
	router.Post("/create", handler.Create)
	router.Put("/:id/edit", handler.Update)
	router.Delete("/:id/delete", handler.Delete)
}
