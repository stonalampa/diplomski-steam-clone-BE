package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type LibraryRecord struct {
	ID          primitive.ObjectID   `bson:"_id"`
	UserId      primitive.ObjectID   `json:"userId" bson:"userId"`
	GameIds     []primitive.ObjectID `json:"gameIds" bson:"gameIds"`
	WishlistIds []primitive.ObjectID `json:"wishlistIds" bson:"wishlistIds"`
}
