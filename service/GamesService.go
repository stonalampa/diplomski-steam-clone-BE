package service

import (
	"main/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GamesService struct {
	gamesRepository repository.GamesRepository
}

func NewGamesService(gamesRepository repository.GamesRepository) *GamesService {
	return &GamesService{gamesRepository: gamesRepository}
}

func (gs GamesService) CreateGame(ctx *gin.Context) {
	var game repository.Game
	ctx.BindJSON(&game)

	game.ID = primitive.NewObjectID()

	insertedGame, err := gs.gamesRepository.CreateGame(ctx, &game)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"game": insertedGame})
}

func (gs GamesService) GetGame(ctx *gin.Context) {
	idParam := ctx.Param("id")
	if idParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id parameter is required"})
		return
	}

	gameID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}

	game, err := gs.gamesRepository.GetGame(ctx, gameID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"game": game})
}

func (gs GamesService) FindGames(ctx *gin.Context) {
	games, err := gs.gamesRepository.GetAllGames(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"games": games})
}

func (gs GamesService) GetAllGames(ctx *gin.Context) {
	games, err := gs.gamesRepository.GetAllGames(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"games": games})
}

func (gs GamesService) UpdateGame(ctx *gin.Context) {
	var game repository.Game
	ctx.BindJSON(&game)

	result, err := gs.gamesRepository.UpdateGame(ctx, &game)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if result.ModifiedCount == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "game not found"})
	}

	ctx.JSON(http.StatusOK, gin.H{"game": game})
}

func (gs GamesService) DeleteGame(ctx *gin.Context) {
	var input struct {
		ID primitive.ObjectID `json:"id"`
	}

	err := ctx.BindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	result, err := gs.gamesRepository.DeleteGame(ctx, input.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if result.DeletedCount == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "game not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "game deleted successfully"})
}
