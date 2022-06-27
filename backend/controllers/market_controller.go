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
	"go.mongodb.org/mongo-driver/bson"
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
		Id:        primitive.NewObjectID(),
		Data:      market.Data,
		Timestamp: primitive.Timestamp{T: uint32(time.Now().Unix())},
	}

	_, err := marketCollection.InsertOne(ctx, newMarketData)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.MarketResponse{Status: http.StatusInternalServerError, Message: "Error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.MarketResponse{Status: http.StatusCreated, Message: "Success", Data: &fiber.Map{"data": newMarketData.Data}})
}

func GetMarketDataByRange(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	init_date := c.Params("init_date")
	end_date := c.Params("end_date")
	var market []models.Market
	defer cancel()

	// TODO
	// [] FILTER BY RANGE
	err := marketCollection.Find(ctx, bson.M{"timestamp": bson.M{"$gte": init_date, "$lte": end_date}}).All(ctx, &market)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.MarketResponse{Status: http.StatusInternalServerError, Message: "Error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.MarketResponse{Status: http.StatusOK, Message: "Success", Data: &fiber.Map{"data": market}})
}
