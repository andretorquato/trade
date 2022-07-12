package routes

import (
	"trade/controllers"

	"github.com/gofiber/fiber/v2"
)

func WalletRouter(app *fiber.App) {
	app.Post("/wallet/create", controllers.CreateWallet)
}
