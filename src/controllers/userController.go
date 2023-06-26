package controllers

import (
	"go-gin-rest-api/src/databases"
	"go-gin-rest-api/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
 * Register User.
 */
func RegisterUser(context *gin.Context) {
	var user models.User

	// Binds an input and maps it into user.
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	// Hash the password.
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	// Stores the user data into the database.
	record := databases.UserDB.Create(&user)

	// Aborts if there is an error while saving the data.
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}

/*
 * Gets all users. For testing purposes.
 */
/*
func AllUsers(c *gin.Context) {
	var users []models.User
	databases.UserDB.Find(&users)
	c.IndentedJSON(http.StatusOK, users)
}
*/
