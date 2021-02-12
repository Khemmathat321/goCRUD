package handler

import "github.com/gofiber/fiber/v2"

// Get func
func Get(c *fiber.Ctx) error {
	id := c.Params("id")

	// if id == "" {
	// 	// Do get All data
	// } else {
	// 	// Get data Id
	// }
	println(id == "")

	return c.SendString(id)
}

// Create func
func Create(c *fiber.Ctx) error {
	return c.SendString("Create")
}

// Update func
func Update(c *fiber.Ctx) error {
	return c.SendString("Update")

}

// Delete func
func Delete(c *fiber.Ctx) error {
	return c.SendString("Delete")
}
