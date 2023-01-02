package database

import (
	models "main/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var users = []models.User{
	{
		ID:          primitive.NewObjectID(),
		Email:       "test1@email.com",
		Username:    "SolidStojan1",
		Password:    "Test123123123",
		Name:        "Stole",
		DateOfBirth: "12-03-200012",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          primitive.NewObjectID(),
		Email:       "test2@email.com",
		Username:    "SolidStojan2",
		Password:    "Test123123123",
		Name:        "Stole",
		DateOfBirth: "12-03-200012",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          primitive.NewObjectID(),
		Email:       "test3@email.com",
		Username:    "SolidStojan3",
		Password:    "Test123123123",
		Name:        "Stole",
		DateOfBirth: "12-03-200012",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}

var Seeds = map[string]any{
	"Users":   users,
	"neUsers": users,
}
