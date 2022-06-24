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

var marketCollection *mongo.Collection = configs.GetCollection(configs.DB, "market")
var validate = validator.New()

func CreateMarketData(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var market models.Market
	defer cancel()

	// validate the request body
	if err := c.BodyParser(&market); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.MarketResponse{Status: http.StatusBadRequest, Message: "Error", Data: &fiber.Map{"data": err.Error()}})
	}

	// use the validator library to validate required fields
	if validationErro := validate.Struct(&market); validationErro != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.MarketResponse{Status: http.StatusBadRequest, Message: "Error", Data: &fiber.Map{"data": validationErro.Error()}})
	}

	newMarketData := models.Market{
		Id:   primitive.NewObjectID(),
		Data: market.Data,
	}

	_, err := marketCollection.InsertOne(ctx, newMarketData)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.MarketResponse{Status: http.StatusInternalServerError, Message: "Error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.MarketResponse{Status: http.StatusCreated, Message: "Success", Data: &fiber.Map{"data": newMarketData.Data}})
}