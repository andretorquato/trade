package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Market struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	Data      string             `json:"data,omitempty" validate:"required"`
	Timestamp time.Time          `json:"timestamp"`
}
