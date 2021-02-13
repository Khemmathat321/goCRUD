package main

import (
	database "github/Khemmathat321/goCRUD/database"
	middleware "github/Khemmathat321/goCRUD/middleware"
	model "github/Khemmathat321/goCRUD/models"
	router "github/Khemmathat321/goCRUD/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Init connection to database
	database.PgConnect()
	model.User{}.Migrate()

	// Register routes
	middleware.Registers(app)
	router.Registers(app)

	app.Listen(":3000")
}
