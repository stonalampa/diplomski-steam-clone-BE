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
	GetReview(ctx context.Context, id primitive.ObjectID) (*mongo.InsertOneResult, error)
	GetAllReviewsForGame(ctx context.Context, gameId primitive.ObjectID) ([]Review, error)
	GetAllReviewsFromUser(ctx context.Context, userId primitive.ObjectID) ([]Review, error)
	DeleteReview(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error)
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

func (repo *reviewsRepository) GetReview(ctx context.Context, id primitive.ObjectID) (*mongo.InsertOneResult, error) {
	result, err := repo.db.Collection("reviews").InsertOne(ctx, id)
	return result, err
}

func (repo *reviewsRepository) GetAllReviewsForGame(ctx context.Context, gameId primitive.ObjectID) ([]Review, error) {
	filter := bson.M{"gameId": gameId}
	cursor, err := repo.db.Collection("reviews").Find(ctx, filter)
	if err != nil {
		return []Review{}, err
	}

	var results []Review
	if err = cursor.All(context.TODO(), &results); err != nil {
		return []Review{}, err
	}
	return results, nil
}

func (repo *reviewsRepository) GetAllReviewsFromUser(ctx context.Context, userId primitive.ObjectID) ([]Review, error) {
	filter := bson.M{"userId": userId}
	cursor, err := repo.db.Collection("reviews").Find(ctx, filter)
	if err != nil {
		return []Review{}, err
	}

	var results []Review
	if err = cursor.All(context.TODO(), &results); err != nil {
		return []Review{}, err
	}
	return results, nil
}

func (repo *reviewsRepository) DeleteReview(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	res, err := repo.db.Collection("reviews").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return &mongo.DeleteResult{}, err
	}
	return res, nil
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
