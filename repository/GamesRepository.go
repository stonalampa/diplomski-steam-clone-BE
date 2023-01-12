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
	GetGame(ctx context.Context, id primitive.ObjectID) (Game, error)
	GetGames(ctx context.Context, numberOfRecords int64) ([]Game, error)
	GetAllGames(ctx context.Context) ([]Game, error)
	UpdateGame(ctx context.Context, game *Game) (*mongo.UpdateResult, error)
	DeleteGame(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error)
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

func (repo *gamesRepository) GetGame(ctx context.Context, id primitive.ObjectID) (Game, error) {
	var game Game
	err := repo.db.Collection("games").FindOne(ctx, bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&game)
	if err != nil {
		return Game{}, err
	}

	return game, nil
}

func (repo *gamesRepository) GetGames(ctx context.Context, numberOfRecords int64) ([]Game, error) {
	opts := options.Find().SetLimit(numberOfRecords)
	cursor, err := repo.db.Collection("games").Find(ctx, bson.D{}, opts)
	if err != nil {
		return []Game{}, err
	}

	var results []Game
	if err = cursor.All(context.TODO(), &results); err != nil {
		return []Game{}, err
	}
	return results, nil
}

func (repo *gamesRepository) GetAllGames(ctx context.Context) ([]Game, error) {
	cursor, err := repo.db.Collection("games").Find(ctx, bson.D{})
	if err != nil {
		return []Game{}, err
	}

	var results []Game
	if err = cursor.All(context.TODO(), &results); err != nil {
		return []Game{}, err
	}
	return results, nil
}

func (repo *gamesRepository) UpdateGame(ctx context.Context, game *Game) (*mongo.UpdateResult, error) {
	res, err := repo.db.Collection("games").UpdateByID(ctx, game.ID, game)
	if err != nil {
		return &mongo.UpdateResult{}, err
	}

	return res, nil
}

func (repo *gamesRepository) DeleteGame(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	res, err := repo.db.Collection("games").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return &mongo.DeleteResult{}, err
	}
	return res, nil
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
