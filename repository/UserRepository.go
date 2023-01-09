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

// podaci o useru
// payment card podaci, to isto enkriptuj i dekriptuj kada vratis iz baze mada i ne mora realno me boli q
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
