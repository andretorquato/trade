package controllers

import (
	"context"
	"net/http"
	"time"
	"trade/configs"
	"trade/models"
	"trade/responses"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var walletCollection *mongo.Collection = configs.GetCollection(configs.DB, "wallet")

func CreateWallet(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var wallet models.Wallet
	defer cancel()

	if err := c.BodyParser(&wallet); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.WalletResponse{Status: http.StatusBadRequest, Message: "Error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErro := validator.New().Struct(&wallet); validationErro != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.WalletResponse{Status: http.StatusBadRequest, Message: "Error", Data: &fiber.Map{"data": validationErro.Error()}})
	}

	newWalletData := models.Wallet{
		Id:     primitive.NewObjectID(),
		Name:   wallet.Name,
		Tokens: append(wallet.Tokens, wallet.Tokens...),
	}

	_, err := walletCollection.InsertOne(ctx, newWalletData)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.WalletResponse{Status: http.StatusInternalServerError, Message: "Error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.WalletResponse{Status: http.StatusOK, Message: "Success", Data: &fiber.Map{"data": wallet}})
}
