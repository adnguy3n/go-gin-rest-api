package controllers

import (
	"go-gin-rest-api/src/authenthication"
	"go-gin-rest-api/src/databases"
	"go-gin-rest-api/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
 * Struct for Token Requests.
 */
type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*
 * Generate Token function.
 */
func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user models.User

	// Bind the incoming request to TokenRequest.
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	// Check if email exists and password is correct.
	record := databases.UserDB.Where("email = ?", request.Email).First(&user)

	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	// Check if the password matches the password in the database.
	credentialError := user.CheckPassword(request.Password)

	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}

	// Generate the JWT.
	tokenString, err := authenthication.GenerateJWT(user.Email, user.Username)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
