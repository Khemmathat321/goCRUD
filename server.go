package main

import (
	middleware "github/Khemmathat321/goCRUD/middleware"
	router "github/Khemmathat321/goCRUD/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// db, _ := database.PgConnect()

	middleware.Registers(app)
	router.Registers(app)

	app.Listen(":3000")
}
