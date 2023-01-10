package repository

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Game struct {
	ID              primitive.ObjectID `bson:"_id"`
	Title           string             `json:"userId" bson:"userId"`
	Price           float32            `json:"gameId" bson:"gameId"`
	Description     string             `json:"comment" bson:"comment"`
	Score           float32            `json:"score" bson:"score"`
	NumberOfScores  int                `json:"numberOfScores" bson:"numberOfScores"`
	Screenshots     []string           `json:"screenshots" bson:"screenshots"`
	Discount        float32            `json:"discount" bson:"discount"`
	DiscountExpDate time.Time          `json:"discountExpDate" bson:"discountExpDate"`
	ReleaseDate     time.Time          `json:"releaseDate" bson:"releaseDate"`
}
