package main

import (
	"trade/configs"
	"trade/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// run database
	configs.ConnectDB()

	routes.MarketRouter(app)
	routes.WalletRouter(app)

	app.Listen(":6000")
}
