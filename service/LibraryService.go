package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"main/helpers"
	"main/repository"
)

type LibraryService struct {
	libraryRepository repository.LibraryRepository
	gamesRepository   repository.GamesRepository
}

func NewLibraryService(libraryRepository repository.LibraryRepository, gamesRepository repository.GamesRepository) *LibraryService {
	return &LibraryService{
		libraryRepository: libraryRepository,
		gamesRepository:   gamesRepository,
	}
}

func (ls *LibraryService) CreateLibraryRecord(ctx *gin.Context) {
	var libraryRecord repository.LibraryRecord
	err := ctx.BindJSON(&libraryRecord)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	libraryRecord.ID = primitive.NewObjectID()

	result, err := ls.libraryRepository.CreateLibraryRecord(ctx, &libraryRecord)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, result.InsertedID)
}

func (ls *LibraryService) GetLibraryRecord(ctx *gin.Context) {
	idParam := ctx.Param("id")
	if idParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id parameter is required"})
		return
	}

	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}

	record, err := ls.libraryRepository.GetLibraryRecord(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	ctx.JSON(http.StatusOK, record)
}

func (ls *LibraryService) UpdateLibraryRecord(ctx *gin.Context) {
	email := ctx.Query("email")
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email parameter is required"})
		return
	}

	removeStr := ctx.Query("remove")
	remove, err := strconv.ParseBool(removeStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid value for remove parameter"})
		return
	}

	var libraryRecord repository.LibraryRecord
	err = ctx.BindJSON(&libraryRecord)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if libraryRecord.ID.IsZero() {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	updateGameIds := make([]primitive.ObjectID, 0)
	updateWishlistIds := make([]primitive.ObjectID, 0)

	// Check if gameIds field is provided in the request
	sendEmail := false
	if len(libraryRecord.GameIds) > 0 {
		updateGameIds = libraryRecord.GameIds
		sendEmail = true
	}

	// Check if wishlistIds field is provided in the request
	if len(libraryRecord.WishlistIds) > 0 {
		updateWishlistIds = libraryRecord.WishlistIds
	}

	// Create a new LibraryRecord with only the ID and UserID
	updatedLibraryRecord := &repository.LibraryRecord{
		ID:          libraryRecord.ID,
		UserId:      libraryRecord.UserId,
		WishlistIds: updateWishlistIds,
		GameIds:     updateGameIds,
	}

	result, err := ls.libraryRepository.UpdateLibraryRecord(ctx, updatedLibraryRecord, remove)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.ModifiedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	if sendEmail {
		game, err := repository.GamesRepository.GetGame(ls.gamesRepository, ctx, libraryRecord.GameIds[0])
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if game.Discount != 0 {
			game.Price = game.Price - game.Discount
		}

		priceToString := strconv.FormatFloat(float64(game.Price), 'f', -1, 32)
		success := helpers.GenerateSuccessfulPurchaseEmail(email, game.Title, priceToString, "emailTemplates/SuccessfulPayment.html")
		if !success {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error sending email"})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "record updated"})
}

func (ls *LibraryService) DeleteLibraryRecord(ctx *gin.Context) {
	idParam := ctx.Param("id")
	if idParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id parameter is required"})
		return
	}

	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}

	result, err := ls.libraryRepository.DeleteLibraryRecord(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "record deleted successfully"})
}
