package seeds

import (
	repo "main/repository"
	"main/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var password, err = utils.HashPassword("Test123123123")
var users = []repo.User{
	{
		ID:          primitive.NewObjectID(),
		Email:       "test1@email.com",
		Username:    "SolidStojan1",
		Password:    password,
		Name:        "Stole",
		DateOfBirth: "31-12-1999",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          primitive.NewObjectID(),
		Email:       "test2@email.com",
		Username:    "SolidStojan2",
		Password:    password,
		Name:        "Stole",
		DateOfBirth: "31-12-1999",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          primitive.NewObjectID(),
		Email:       "test3@email.com",
		Username:    "SolidStojan3",
		Password:    password,
		Name:        "Stole",
		DateOfBirth: "31-12-1999",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}

var Seeds = map[string]any{
	"Users": users,
}
