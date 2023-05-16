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

type UsersRepository interface {
	CreateUser(ctx context.Context, user *User) (*mongo.InsertOneResult, error)
	GetUser(ctx context.Context, id primitive.ObjectID) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetAllUsers(ctx context.Context) ([]User, error)
	UpdateUser(ctx context.Context, data User) (*mongo.UpdateResult, error)
	DeleteUser(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error)
	DropUsers(ctx context.Context)
	CreateIndices(ctx context.Context)
}

type usersRepository struct {
	db *mongo.Database
}

func NewUsersRepository(db *mongo.Database) UsersRepository {
	return &usersRepository{db: db}
}

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	Username     string             `json:"username" bson:"username"`
	Email        string             `json:"email" bson:"email"`
	Password     string             `json:"password" bson:"password"`
	Name         string             `json:"name" bson:"name"`
	DateOfBirth  string             `json:"dateOfBirth" bson:"dateOfBirth"`
	IsAdmin      bool               `json:"isAdmin" bson:"isAdmin"`
	IsActive     bool               `json:"isActive" bson:"isActive"`
	PaymentCards []PaymentCard      `json:"paymentCards" bson:"paymentCards"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
}

type PaymentCard struct {
	CardNumber string
	ExpDate    string
	Cvc        int
}

func (repo *usersRepository) CreateUser(ctx context.Context, user *User) (*mongo.InsertOneResult, error) {
	result, err := repo.db.Collection("users").InsertOne(ctx, user)
	return result, err
}

func (repo *usersRepository) GetUser(ctx context.Context, id primitive.ObjectID) (User, error) {
	var user User
	err := repo.db.Collection("users").FindOne(ctx, bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (repo *usersRepository) GetUserByEmail(ctx context.Context, email string) (User, error) {
	var user User
	err := repo.db.Collection("users").FindOne(ctx, bson.D{primitive.E{Key: "email", Value: email}}).Decode(&user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (repo *usersRepository) GetAllUsers(ctx context.Context) ([]User, error) {
	cursor, err := repo.db.Collection("users").Find(ctx, bson.D{})
	if err != nil {
		return []User{}, err
	}

	var results []User
	if err = cursor.All(context.TODO(), &results); err != nil {
		return []User{}, err
	}
	return results, nil
}

func (repo *usersRepository) UpdateUser(ctx context.Context, data User) (*mongo.UpdateResult, error) {
	update := bson.M{
		"$set": data,
	}

	res, err := repo.db.Collection("users").UpdateByID(ctx, data.ID, update)
	if err != nil {
		return &mongo.UpdateResult{}, err
	}

	return res, nil
}

func (repo *usersRepository) DeleteUser(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	res, err := repo.db.Collection("users").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return &mongo.DeleteResult{}, err
	}
	return res, nil
}

func (repo *usersRepository) DropUsers(ctx context.Context) {
	_, err := repo.db.Collection("users").DeleteMany(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
}

func (repo *usersRepository) CreateIndices(ctx context.Context) {
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
