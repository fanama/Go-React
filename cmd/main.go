package main

import (
	"fmt"

	"github.com/fanama/Go-React/cmd/variables"
	"github.com/fanama/Go-React/package/middleware"
	"github.com/fanama/Go-React/package/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	api := app.Group("/api", cors.New())

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("root of the api")
	})

	api.Get("/name", middleware.GetName)

	app.Static("/", "./public")

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/login", router.Login)

	port := fmt.Sprintf(":%d", variables.Config.ServerPort)

	app.Listen(port)
}
