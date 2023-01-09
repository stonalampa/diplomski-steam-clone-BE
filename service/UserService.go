package service

import (
	"main/repository"
	"main/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *Service {
	return &Service{userRepository: userRepository}
}

func (s Service) CreateUser(ctx *gin.Context) {
	validated := validateCreateUser(ctx)
	if !validated {
		return
	}

	var user repository.User
	ctx.BindJSON(&user)

	hashedPass, err := utils.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Password = hashedPass

	insertedUser, err := s.userRepository.CreateUser(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": insertedUser})
}

func (s Service) GetUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"hello": "world"})
}
