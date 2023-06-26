package controllers

import "github.com/gin-gonic/gin"

/*
 * HomePage endpoint. Prints out a message when hit.
 */
func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"Oh hi there.": "This is Home.",
	})
}

/*
 * HomePage endpoint for POST. Only hits if a POST request is made instead of a GET request.
 * Written for learning purposes.
 */
func HomePagePOST(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Silly, why did you use a POST request to get Home?",
	})
}
