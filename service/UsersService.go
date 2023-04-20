package service

import (
	"main/repository"
	"main/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	usersRepository repository.UsersRepository
}

func NewUsersService(usersRepository repository.UsersRepository) *UserService {
	return &UserService{usersRepository: usersRepository}
}

func (us UserService) GetUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := us.usersRepository.GetUser(ctx, objID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (us UserService) GetUsers(ctx *gin.Context) {
	users, err := us.usersRepository.GetAllUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func (us UserService) CreateUser(ctx *gin.Context) {
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

	insertedUser, err := us.usersRepository.CreateUser(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": insertedUser})
}

func (us UserService) UpdateUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var updatedUser repository.User
	ctx.BindJSON(&updatedUser)

	hashedPass, err := utils.HashPassword(updatedUser.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	updatedUser.ID, err = primitive.ObjectIDFromHex(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	updatedUser.UpdatedAt = time.Now()
	updatedUser.Password = hashedPass

	_, err = us.usersRepository.UpdateUser(ctx, updatedUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (us UserService) DeleteUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	result, err := us.usersRepository.DeleteUser(ctx, objID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if result.DeletedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
