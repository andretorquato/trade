package routes

import "github.com/gofiber/fiber/v2"

func UserRoute(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Init Trader backend"})
	})
}
