package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"main/repository"
)

type LibraryService struct {
	libraryRepository repository.LibraryRepository
}

func NewLibraryService(libraryRepository repository.LibraryRepository) *LibraryService {
	return &LibraryService{libraryRepository: libraryRepository}
}

func (ls *LibraryService) CreateLibraryRecord(ctx *gin.Context) {
	var libraryRecord repository.LibraryRecord
	err := ctx.BindJSON(&libraryRecord)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Assign a new ObjectID to the record
	libraryRecord.ID = primitive.NewObjectID()

	// Insert the record into the database
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
	var libraryRecord repository.LibraryRecord
	err := ctx.BindJSON(&libraryRecord)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Ensure that the record ID is valid
	if libraryRecord.ID.IsZero() {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	// Update the record in the database
	result, err := ls.libraryRepository.UpdateLibraryRecord(ctx, libraryRecord)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.ModifiedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
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

	// Delete the record from the database
	result, err := ls.libraryRepository.DeleteLibraryRecord(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	// Return success message if the record was deleted
	ctx.JSON(http.StatusOK, gin.H{"message": "record deleted successfully"})
}
