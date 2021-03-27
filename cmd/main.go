package main

import (
	"github.com/fanama/Go-React/package/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/", "./public")

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/login", router.Login)

	app.Listen(":3000")
}
