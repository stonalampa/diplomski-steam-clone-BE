package service

import (
	"main/repository"
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthService struct {
	usersRepository repository.UsersRepository
}

func NewAuthService(usersRepository repository.UsersRepository) *AuthService {
	return &AuthService{usersRepository: usersRepository}
}

func (as AuthService) Login(ctx *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := ctx.BindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	user, err := as.usersRepository.GetUserByEmail(ctx, input.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user.IsAdmin {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Admin user trying to login as a normal user"})
		return
	}

	if !user.IsActive {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User is not active"})
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.Email, true, user.IsAdmin)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (as AuthService) AdminLogin(ctx *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := ctx.BindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	user, err := as.usersRepository.GetUserByEmail(ctx, input.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !user.IsAdmin || !utils.CheckPasswordHash(input.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(input.Email, true, user.IsAdmin)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
