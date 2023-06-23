package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
 * All Items endpoint. Returns a json response of all Items when hit.
 */
func allItems(c *gin.Context) {
	fmt.Println("Hit all Items endpoint")
	c.IndentedJSON(http.StatusOK, Items)
}

/*
 * Appends an item from JSON received in the request body.
 */
func postItem(c *gin.Context) {
	var newItem Item

	// Call BindJSON to bind the received JSON to newItem.
	if err := c.BindJSON(&newItem); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "input not found"})
		return
	}

	Items = append(Items, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)
}

/*
 * Gets the item whose ID value matches the id given.
 */
func getItem(c *gin.Context) {
	id := c.Param("id")

	// Loops over the list of Items to find an item with a matching ID value.
	for _, i := range Items {
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
func deleteItem(c *gin.Context) {
	id := c.Param("id")

	// Loops over the list of Items to find an item with a matching ID value.
	for index, i := range Items {
		if i.ID == id {
			// Delete Item from slice
			Items = append(Items[:index], Items[index+1:]...)

			c.IndentedJSON(http.StatusOK, i)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

/*
 * Patches the item whose ID value matches the id given.
 */
func patchItem(c *gin.Context) {
	id := c.Param("id")

	var updatedItem Item

	// Call BindJSON to bind the received JSON to newItem.
	if err := c.BindJSON(&updatedItem); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "input not found"})
		return
	}

	// Loops over the list of Items to find an item with a matching ID value.
	for index := range Items {
		if Items[index].ID == id {
			if updatedItem.ID != "" {
				Items[index].ID = updatedItem.ID
			}

			if updatedItem.Name != "" {
				Items[index].Name = updatedItem.Name
			}

			if updatedItem.Desc != "" {
				Items[index].Desc = updatedItem.Desc
			}

			if updatedItem.Content != "" {
				Items[index].Content = updatedItem.Content
			}

			c.IndentedJSON(http.StatusOK, Items[index])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

/*
 * PUT the item whose ID value matches the id given.
 */
func putItem(c *gin.Context) {
	id := c.Param("id")

	var updatedItem Item

	// Call BindJSON to bind the received JSON to newItem.
	if err := c.BindJSON(&updatedItem); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "input not found"})
		return
	}

	// Loops over the list of Items to find an item with a matching ID value.
	for index := range Items {
		if Items[index].ID == id {
			Items[index] = updatedItem
			c.IndentedJSON(http.StatusOK, Items[index])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

/*
 * HomePage endpoint. Prints out a message when hit.
 */
func homePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hi, this is home",
	})
}

/*
 * HomePage endpoint for POST. Only hits if a POST request is made instead of a GET request.
 */
func homePagePOST(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hi, this is POST home",
	})
}
