package controllers

import (
	"fmt"
	"go-gin-rest-api/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
 * All items endpoint. Returns a json response of all items when hit.
 */
func AllItems(c *gin.Context) {
	fmt.Println("Hit all Items endpoint")
	c.IndentedJSON(http.StatusOK, models.Items)
}

/*
 * Appends an item from JSON received in the request body.
 */
func PostItem(c *gin.Context) {
	var newItem models.Item

	// Call BindJSON to bind the received JSON to newItem.
	if err := c.BindJSON(&newItem); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "input not found"})
		return
	}

	models.Items = append(models.Items, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)
}

/*
 * Gets the item whose ID value matches the id given.
 */
func GetItem(c *gin.Context) {
	id := c.Param("id")

	// Loops over the list of items to find an item with a matching ID value.
	for _, i := range models.Items {
		if i.ID == id {
			c.IndentedJSON(http.StatusOK, i)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

/*
 * Deletes the item whose ID value matches the id given.
 */
func DeleteItem(c *gin.Context) {
	id := c.Param("id")

	// Loops over the list of items to find an item with a matching ID value.
	for index, i := range models.Items {
		if i.ID == id {
			// Delete Item from slice
			models.Items = append(models.Items[:index], models.Items[index+1:]...)

			c.IndentedJSON(http.StatusOK, i)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

/*
 * Patches the item whose ID value matches the id given.
 */
func PatchItem(c *gin.Context) {
	id := c.Param("id")

	var updatedItem models.Item

	// Call BindJSON to bind the received JSON to newItem.
	if err := c.BindJSON(&updatedItem); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "input not found"})
		return
	}

	// Loops over the list of items to find an item with a matching ID value.
	for index := range models.Items {
		if models.Items[index].ID == id {
			if updatedItem.ID != "" {
				models.Items[index].ID = updatedItem.ID
			}

			if updatedItem.Name != "" {
				models.Items[index].Name = updatedItem.Name
			}

			if updatedItem.Desc != "" {
				models.Items[index].Desc = updatedItem.Desc
			}

			if updatedItem.Content != "" {
				models.Items[index].Content = updatedItem.Content
			}

			c.IndentedJSON(http.StatusOK, models.Items[index])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

/*
 * PUT the item whose ID value matches the id given.
 */
func PutItem(c *gin.Context) {
	id := c.Param("id")

	var updatedItem models.Item

	// Call BindJSON to bind the received JSON to newItem.
	if err := c.BindJSON(&updatedItem); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "input not found"})
		return
	}

	// Loops over the list of items to find an item with a matching ID value.
	for index := range models.Items {
		if models.Items[index].ID == id {
			models.Items[index] = updatedItem
			c.IndentedJSON(http.StatusOK, models.Items[index])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

/*
 * HomePage endpoint. Prints out a message when hit.
 */
func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hi, this is home.",
	})
}

/*
 * HomePage endpoint for POST. Only hits if a POST request is made instead of a GET request.
 * Written for learning purposes.
 */
func HomePagePOST(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hi, this is POST home.",
	})
}
