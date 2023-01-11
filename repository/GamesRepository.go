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

type GamesRepository interface {
	CreateGame(ctx context.Context, game *Game) (*mongo.InsertOneResult, error)
	DropGames(ctx context.Context)
	CreateIndices(ctx context.Context)
}

type gamesRepository struct {
	db *mongo.Database
}

func NewGamesRepository(db *mongo.Database) GamesRepository {
	return &gamesRepository{db: db}
}

type Game struct {
	ID              primitive.ObjectID `bson:"_id"`
	Title           string             `json:"title" bson:"title"`
	Price           float32            `json:"price" bson:"price"`
	Developer       string             `json:"developer" bson:"developer"`
	Publisher       string             `json:"publisher" bson:"publisher"`
	Description     string             `json:"description" bson:"description"`
	Score           int                `json:"score" bson:"score"`
	NumberOfScores  int                `json:"numberOfScores" bson:"numberOfScores"`
	Screenshots     []string           `json:"screenshots" bson:"screenshots"`
	Discount        float32            `json:"discount" bson:"discount"`
	DiscountExpDate time.Time          `json:"discountExpDate" bson:"discountExpDate"`
	ReleaseDate     time.Time          `json:"releaseDate" bson:"releaseDate"`
}

func (repo *gamesRepository) CreateGame(ctx context.Context, game *Game) (*mongo.InsertOneResult, error) {
	result, err := repo.db.Collection("games").InsertOne(ctx, game)
	return result, err
}

func (repo *gamesRepository) DropGames(ctx context.Context) {
	_, err := repo.db.Collection("games").DeleteMany(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
}

func (repo *gamesRepository) CreateIndices(ctx context.Context) {
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "title", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	name, err := repo.db.Collection("games").Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
	}
	fmt.Println("Name of Index Created: " + name)
}
