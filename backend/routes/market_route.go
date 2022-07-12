package routes

import (
	"trade/controllers"

	"github.com/gofiber/fiber/v2"
)

func MarketRouter(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "online"})
	})

	app.Post("/new-market-data", controllers.CreateMarketData)
	app.Post("/filter", controllers.GetMarketDataByRange)
}
