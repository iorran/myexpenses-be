package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type InvoiceRequestBody struct {
	Date        string
	Image       string
	Description string
}

type InvoiceDocument struct {
	Id          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Date        time.Time          `json:"date"`
	Image       string             `json:"image"`
	Description string             `json:"description"`
}
