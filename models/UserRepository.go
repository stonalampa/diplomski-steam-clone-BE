package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type M map[string]interface{}

// Define our errors:
var internalError = M{"message": "internal error"}
var userNotFound = M{"message": "user not found"}

type UserRepository interface {
	CreateUser(ctx context.Context, user User) (*mongo.InsertOneResult, error)
	CreateNeUser(ctx context.Context, user User) (*mongo.InsertOneResult, error)
}

type userRepository struct {
	client *mongo.Client
}

type User struct {
	ID          primitive.ObjectID `bson:"_id"`
	Username    string             `json:"username" bson:"username"`
	Email       string             `json:"email" bson:"email"`
	Password    string             `json:"password" bson:"password"`
	Name        string             `json:"name" bson:"name"`
	DateOfBirth string             `json:"dateOfBirth" bson:"dateOfBirth"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

func (repo *userRepository) CreateUser(ctx context.Context, user *User) (*mongo.InsertOneResult, error) {
	res, err := repo.client.Database("loyalty-be-db").Collection("users").InsertOne(ctx, user)
	if err != nil {
		fmt.Println("Error while inserting User", err)
	}
	return res, err
}
