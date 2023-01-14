package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

func GenerateId(hex string) primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex(hex)
	return id
}

func GeneratePassword(pass string) string {
	hashedPass, _ := HashPassword(pass)
	return hashedPass
}
