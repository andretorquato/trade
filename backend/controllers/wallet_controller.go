package controllers

import (
	"context"
	"net/http"
	"time"
	"trade/models"
	"trade/responses"

	"github.com/gofiber/fiber/v2"
)

// var walletCollection *mongo.Collection = configs.GetCollection(configs.DB, "wallet")

func CreateWallet(c *fiber.Ctx) error {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var wallet models.Wallet
	defer cancel()

	// if err := c.BodyParser(&wallet); err != nil {
	// 	return c.Status(http.StatusBadRequest).JSON(responses.WalletResponse{Status: http.StatusBadRequest, Message: "Error", Data: &fiber.Map{"data": err.Error()}})
	// }
	return c.Status(http.StatusOK).JSON(responses.WalletResponse{Status: http.StatusOK, Message: "Success", Data: &fiber.Map{"data": wallet}})
}
