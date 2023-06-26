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
func RegisterUser(c *gin.Context) {
	var user models.User

	// Binds an input and maps it into user.
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	// Hash the password.
	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	// Stores the user data into the database.
	record := databases.UserDB.Create(&user)

	// Aborts if there is an error while saving the data.
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}

/*
 * Gets the user whose username value matches the username given.
 */
func GetUser(c *gin.Context) {
	var user models.User

	// Finds the user based on their unique username.
	// Gives an error if no user with that username exists.
	err := databases.UserDB.Where("username = ?", c.Param("username")).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Username not found."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username, "name": user.Name})
}

/*
 * Update the user information.
 */
func UpdateUser(c *gin.Context) {
	var user models.User
	var updatedUser models.User

	// Finds the user based on their unique username.
	// Gives an error if no user with that username exists.
	err := databases.UserDB.Where("username = ?", c.Param("username")).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Username not found."})
		return
	}

	// Binds an input and maps it into updatedUser.
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	// Hash the password.
	if err := updatedUser.HashPassword(updatedUser.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	databases.UserDB.Model(&user).Updates(updatedUser)
	c.IndentedJSON(http.StatusOK, gin.H{"Message": "User Info Updated."})
}

/*
 * Delete User.
 */
func DeleteUser(c *gin.Context) {
	var user models.User

	// Finds the user based on their unique username.
	// Gives an error if no user with that username exists.
	err := databases.UserDB.Where("username = ?", c.Param("username")).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Username not found."})
		return
	}

	//Hard Delete
	databases.UserDB.Unscoped().Delete(&user)
	c.JSON(http.StatusOK, gin.H{"Message": "User deleted."})
}
