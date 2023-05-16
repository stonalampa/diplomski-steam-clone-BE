package utils

import (
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateId(hex string) primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex(hex)
	return id
}

func GeneratePassword(pass string) string {
	hashedPass, _ := HashPassword(pass)
	return hashedPass
}

func GenerateRandomPassword() string {
	rand.Seed(time.Now().UnixNano())

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	password := make([]byte, 8)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	return string(password)
}
