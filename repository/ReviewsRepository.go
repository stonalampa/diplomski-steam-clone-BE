package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type Review struct {
	ID      primitive.ObjectID `bson:"_id"`
	UserId  primitive.ObjectID `json:"userId" bson:"userId"`
	GameId  primitive.ObjectID `json:"gameId" bson:"gameId"`
	Comment string             `json:"comment" bson:"comment"`
}
