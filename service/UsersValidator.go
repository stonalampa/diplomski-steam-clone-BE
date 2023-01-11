package service

import (
	"main/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func validateCreateUser(ctx *gin.Context) bool {
	var user repository.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return false
	}

	return true
}