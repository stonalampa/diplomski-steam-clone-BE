package utils

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type UnsignedResponse struct {
	Message interface{} `json:"message"`
}
type authCustomClaims struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	LoggedIn bool   `json:"loggedIn"`
	IsAdmin  bool   `json:"isAdmin"`
	jwt.StandardClaims
}

type authConfirmationClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var secret = "ReallySecureSecret"
var issuer = "SolidStojan"

func ValidateJwt(c *gin.Context) {
	jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	token, err := parseToken(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "Bad JWT token",
		})
		return
	}

	_, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, UnsignedResponse{
			Message: "Unable to parse claims",
		})
		return
	}
	c.Next()
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("Bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("Incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func parseToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, errors.New("Bad signed method received")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, errors.New("Bad jwt token")
	}

	return token, nil
}

func GenerateToken(id string, email string, loggedIn bool, isAdmin bool) (string, error) {
	// * If development env generate a token that won't expire otherwise it is valid for 2 days
	env := viper.GetString("env")
	var expiresAt int64
	if env == "local" {
		expiresAt = time.Now().Add(time.Hour * 999999).Unix()
	} else {
		expiresAt = time.Now().Add(time.Hour * 48).Unix()
	}

	claims := &authCustomClaims{
		id,
		email,
		loggedIn,
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}
	return signedToken, nil
}

func GenerateConfirmationToken(email string) (string, error) {
	claims := &authConfirmationClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}
	return signedToken, nil
}

func ValidateConfirmationJwt(c *gin.Context) (string, error) {
	tokenString := c.Query("token")
	if tokenString == "" {
		return "", errors.New("invalid token")
	}
	token, err := parseToken(tokenString)
	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return "", errors.New("invalid email claim")
	}

	return email, nil
}
