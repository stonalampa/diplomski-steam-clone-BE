package repository

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (*mongo.InsertOneResult, error)
	DropUsers(ctx context.Context)
	CreateIndices(ctx context.Context)
}

type userRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{db: db}
}

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	Username     string             `json:"username" bson:"username"`
	Email        string             `json:"email" bson:"email"`
	Password     string             `json:"password" bson:"password"`
	Name         string             `json:"name" bson:"name"`
	DateOfBirth  string             `json:"dateOfBirth" bson:"dateOfBirth"`
	IsAdmin      bool               `json:"isAdmin" bson:"isAdmin"`
	PaymentCards []PaymentCard      `json:"paymentCards" bson:"paymentCards"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
}

type PaymentCard struct {
	CardNumber string
	ExpDate    string
	Cvc        int
}

func (repo *userRepository) CreateUser(ctx context.Context, user *User) (*mongo.InsertOneResult, error) {
	result, err := repo.db.Collection("users").InsertOne(ctx, user)
	return result, err
}

// func (repo *userRepository) FindUsers(ctx context.Context) ([]User, error) {

// }

// func (repo *userRepository) GetUser(ctx context.Context, email string) (User, error) {

// }

// func (repo *userRepository) UpdateUser(ctx context.Context, data User) (*mongo.UpdateOneModel, error) {

// }

// func (repo *userRepository) DeleteUser(ctx context.Context, email) (*mongo.DeleteResult, error) {

// }

func (repo *userRepository) DropUsers(ctx context.Context) {
	_, err := repo.db.Collection("users").DeleteMany(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
}

func (repo *userRepository) CreateIndices(ctx context.Context) {
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	name, err := repo.db.Collection("users").Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
	}
	fmt.Println("Name of Index Created: " + name)
}
