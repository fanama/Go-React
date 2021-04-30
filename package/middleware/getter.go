package middleware

import "github.com/gofiber/fiber/v2"

func GetName(c *fiber.Ctx) error {
	return c.SendString("here is your name")
}
