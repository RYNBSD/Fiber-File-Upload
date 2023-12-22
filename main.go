package main

import (
	// "fmt"
	"log"
	"os"
	"path"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	publicDir := path.Join(rootDir(), "public")

	app.Use(recover.New())
	app.Static("/", publicDir)

	app.Post("/file", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			panic(err)
		}

		if err := c.SaveFile(file, path.Join(publicDir, file.Filename)); err != nil {
			panic(err)
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"success": true,
		})
	})

	app.Post("/files", func(c *fiber.Ctx) error {
		form, err := c.MultipartForm()
		if err != nil {
			panic(err)
		}

		files := form.File["files"]
		for _, file := range files {
			if err := c.SaveFile(file, path.Join(publicDir, file.Filename)); err != nil {
				panic(err)
			}
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"success": true,
		})
	})

	log.Fatal(app.Listen(":3000"))
}

func rootDir() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}
