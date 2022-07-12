package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Wallet struct {
	Id     primitive.ObjectID `json:"id,omitempty"`
	Name   string             `json:"name,omitempty" validate:"required"`
	Tokens []Token            `json:"tokens,omitempty"`
}

type Token struct {
	Id     primitive.ObjectID `json:"id,omitempty"`
	Name   string             `json:"name,omitempty" validate:"required"`
	Amount float64            `json:"amount,omitempty" validate:"required"`
}
