package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Market struct {
	Id   primitive.ObjectID `json:"id,omitempty"`
	Data string             `json:"data,omitempty" validate:"required"`
}
