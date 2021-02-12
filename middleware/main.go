package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Registers func
func Registers(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		if form, err := c.MultipartForm(); err == nil {
			// => *multipart.Form

			// Get all files from "documents" key:
			files := form.File["file"]
			// => []*multipart.FileHeader

			// Loop through files:
			for _, file := range files {
				fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
				// => "tutorial.pdf" 360641 "application/pdf"

				// Save the files to disk:
				if err := c.SaveFile(file, fmt.Sprintf("./%s", file.Filename)); err != nil {
					return err
				}
			}
			return err
		}

		return c.Next()
	})
}
