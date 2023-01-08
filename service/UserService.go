package service

import (
	"main/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Server struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *Server {
	return &Server{userRepository: userRepository}
}

func (s Server) CreateUser(ctx *gin.Context) {
	var user repository.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	insertedUser := s.userRepository.CreateUser(ctx, &user)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, err.Error())
	// 	return
	// }
	ctx.JSON(http.StatusOK, gin.H{"user": insertedUser})
}
