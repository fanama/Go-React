package router

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {

	body := c.Body()

	var formulaire form

	err := json.Unmarshal(body, &formulaire)

	if err != nil {
		log.Fatal(err)
	}

	result := fmt.Sprintln(formulaire.Username + " tried to login ...")

	return c.JSON(result)

}
