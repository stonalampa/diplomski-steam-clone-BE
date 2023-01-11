package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ReviewsRepository interface {
	CreateReview(ctx context.Context, review *Review) (*mongo.InsertOneResult, error)
	DropReviews(ctx context.Context)
	CreateIndices(ctx context.Context)
}

type reviewsRepository struct {
	db *mongo.Database
}

func NewReviewsRepository(db *mongo.Database) ReviewsRepository {
	return &reviewsRepository{db: db}
}

type Review struct {
	ID      primitive.ObjectID `bson:"_id"`
	UserId  primitive.ObjectID `json:"userId" bson:"userId"`
	GameId  primitive.ObjectID `json:"gameId" bson:"gameId"`
	Comment string             `json:"comment" bson:"comment"`
}

func (repo *reviewsRepository) CreateReview(ctx context.Context, review *Review) (*mongo.InsertOneResult, error) {
	result, err := repo.db.Collection("reviews").InsertOne(ctx, review)
	return result, err
}

func (repo *reviewsRepository) DropReviews(ctx context.Context) {
	_, err := repo.db.Collection("reviews").DeleteMany(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
}

func (repo *reviewsRepository) CreateIndices(ctx context.Context) {
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "userId", Value: 1}, {Key: "gameId", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	name, err := repo.db.Collection("reviews").Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
	}
	fmt.Println("Name of Index Created: " + name)
}
