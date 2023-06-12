package service

import (
	"main/helpers"
	"main/repository"
	"main/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	usersRepository   repository.UsersRepository
	libraryRepository repository.LibraryRepository
}

func NewUsersService(usersRepository repository.UsersRepository, libraryRepository repository.LibraryRepository) *UserService {
	return &UserService{
		usersRepository:   usersRepository,
		libraryRepository: libraryRepository,
	}
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
	user.IsAdmin = false
	user.IsActive = false

	insertedUser, err := us.usersRepository.CreateUser(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	libraryRecord := &repository.LibraryRecord{
		ID:          primitive.NewObjectID(),
		UserId:      insertedUser.InsertedID.(primitive.ObjectID),
		GameIds:     []primitive.ObjectID{},
		WishlistIds: []primitive.ObjectID{},
	}
	_, err = us.libraryRepository.CreateLibraryRecord(ctx, libraryRecord)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	result := helpers.ConfirmationEmail("test@test.com", "emailTemplates/ConfirmAccountEmail.html")
	if !result {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Email not sent"})
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

func (us UserService) ResetEmail(ctx *gin.Context) {
	var reqBody struct {
		Email string
	}

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := us.usersRepository.GetUserByEmail(ctx, reqBody.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
		return
	}

	newPassword := utils.GenerateRandomPassword()
	user.Password = utils.GeneratePassword(newPassword)

	_, err = us.usersRepository.UpdateUser(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User update failed"})
		return
	}

	result := helpers.GenerateNewPasswordEmail(user.Email, newPassword, "emailTemplates/PasswordResetEmail.html")
	if !result {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User email sending failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User password was reset successfully"})
}

func (us UserService) ConfirmUser(ctx *gin.Context) {
	email, err := utils.ValidateConfirmationJwt(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid confirmation token"})
		return
	}

	user, err := us.usersRepository.GetUserByEmail(ctx, email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
		return
	}

	user.IsActive = true
	_, err = us.usersRepository.UpdateUser(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User update failed"})
		return
	}

	result := helpers.ConfirmationEmail(email, "emailTemplates/ConfirmAccountEmail.html")
	if !result {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User update failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Account successfully activated"})
}
