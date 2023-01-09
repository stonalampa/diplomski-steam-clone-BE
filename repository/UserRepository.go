package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (*mongo.InsertOneResult, error)
	DropUsers(ctx context.Context) bool
}

type userRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{db: db}
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
	result, err := repo.db.Collection("users").InsertOne(ctx, user)
	return result, err
}

func (repo *userRepository) DropUsers(ctx context.Context) bool {
	_, err := repo.db.Collection("users").DeleteMany(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	if r := recover(); r != nil {
		fmt.Println("Recovered!")
	}
	return true
}
