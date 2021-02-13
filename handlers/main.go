package handler

import (
	"github/Khemmathat321/goCRUD/database"
	model "github/Khemmathat321/goCRUD/models"

	"github.com/gofiber/fiber/v2"
)

// Get func
func Get(c *fiber.Ctx) error {
	id := c.Params("id")

	var user model.User
	var users []model.User

	if id != "" {
		// Do get All data
		database.DBConn.First(&user, id)
		return c.JSON(user)
	} else {
		// Get data Id
		database.DBConn.Find(&users)
		return c.JSON(users)
	}
}

// Create func
func Create(c *fiber.Ctx) error {
	data := new(model.User)

	if err := c.BodyParser(data); err != nil {
		return err
	}

	database.DBConn.Create(&data)

	return c.SendString("Created")
}

// Update func
func Update(c *fiber.Ctx) error {
	var user model.User

	database.DBConn.First(&user, c.Params("id"))

	// Pass body value to existing record
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DBConn.Save(&user)

	return c.SendString("Updated")

}

// Delete func
func Delete(c *fiber.Ctx) error {
	database.DBConn.Delete(&model.User{}, c.Params("id"))

	return c.SendString("Deleted")
}
