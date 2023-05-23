package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"main/repository"
)

type ReviewsService struct {
	reviewsRepository repository.ReviewsRepository
}

func NewReviewsService(reviewsRepository repository.ReviewsRepository) *ReviewsService {
	return &ReviewsService{reviewsRepository: reviewsRepository}
}

func (rs *ReviewsService) CreateReviewRecord(ctx *gin.Context) {
	var reviewRecord repository.Review
	err := ctx.BindJSON(&reviewRecord)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	reviewRecord.ID = primitive.NewObjectID()

	result, err := rs.reviewsRepository.CreateReview(ctx, &reviewRecord)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, result.InsertedID)
}

func (rs *ReviewsService) GetReviewRecord(ctx *gin.Context) {
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

	record, err := rs.reviewsRepository.GetReview(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	ctx.JSON(http.StatusOK, record)
}

func (rs *ReviewsService) UpdateReviewRecord(ctx *gin.Context) {
	var reviewRecord repository.Review
	err := ctx.BindJSON(&reviewRecord)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if reviewRecord.ID.IsZero() {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	result, err := rs.reviewsRepository.UpdateReview(ctx, &reviewRecord)
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

func (rs *ReviewsService) DeleteReviewRecord(ctx *gin.Context) {
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

	result, err := rs.reviewsRepository.DeleteReview(ctx, id)
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
