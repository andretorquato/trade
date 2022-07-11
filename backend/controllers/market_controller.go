package controllers

import (
	"context"
	"fmt"
	"log"
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

func CreateMarketData(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var market models.Market
	defer cancel()

	// validate the request body
	if err := c.BodyParser(&market); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.MarketResponse{Status: http.StatusBadRequest, Message: "Error", Data: &fiber.Map{"data": err.Error()}})
	}

	// use the validator library to validate required fields
	if validationErro := validator.New().Struct(&market); validationErro != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.MarketResponse{Status: http.StatusBadRequest, Message: "Error", Data: &fiber.Map{"data": validationErro.Error()}})
	}

	newMarketData := models.Market{
		Id:        primitive.NewObjectID(),
		Data:      market.Data,
		Timestamp: time.Now().UTC(),
	}

	_, err := marketCollection.InsertOne(ctx, newMarketData)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.MarketResponse{Status: http.StatusInternalServerError, Message: "Error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.MarketResponse{Status: http.StatusCreated, Message: "Success", Data: &fiber.Map{"data": newMarketData.Data}})
}

type Range struct {
	InitialDate string `json:"init_date" xml:"init_date" form:"init_date"`
	EndDate     string `json:"end_date" xml:"end_date" form:"end_date"`
}

func GetMarketDataByRange(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	r := new(Range)
	defer cancel()

	if err := c.BodyParser(r); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.MarketResponse{Status: http.StatusBadRequest, Message: "Error", Data: &fiber.Map{"data": err.Error()}})
	}

	const (
		layoutISO = "2006-01-02T15:04:05.000Z"
	)

	formattedInitialDate, _ := time.Parse(layoutISO, r.InitialDate)
	formattedEndDate, _ := time.Parse(layoutISO, r.EndDate)
	fmt.Printf("%v\n", formattedInitialDate)
	fmt.Printf("%v\n", formattedEndDate)

	filter := bson.M{
		"timestamp": bson.M{
			"$gt": formattedInitialDate,
			"$lt": formattedEndDate,
		},
	}

	res, err := marketCollection.Find(ctx, filter)

	var result []bson.M
	if err = res.All(context.TODO(), &result); err != nil {
		log.Fatal(err)
	}

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.MarketResponse{Status: http.StatusInternalServerError, Message: "Error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.MarketResponse{Status: http.StatusOK, Message: "Success", Data: &fiber.Map{"data": result}})
}
