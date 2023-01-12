package service

import (
	"main/repository"
	"net/http"
	"strconv"

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
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"game": insertedGame})
}

func (gs GamesService) GetGame(ctx *gin.Context) {
	param, err := strconv.ParseInt(ctx.Query("records"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := gs.gamesRepository.GetGames(ctx, param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, res)
}

func (gs GamesService) FindGames(ctx *gin.Context) {

}

func (gs GamesService) UpdateGame(ctx *gin.Context) {

}

func (gs GamesService) DeleteGame(ctx *gin.Context) {
	//validation
	var input primitive.ObjectID
	ctx.BindJSON(&input)

}
