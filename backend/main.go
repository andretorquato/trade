package main

import (
	"trade/configs"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// run database
	configs.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Init Trader backend"})
	})

	app.Listen(":6000")
}
