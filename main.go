package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

/*
 * Item data structure.
 */
type Item struct {
	ID      string `json:"id"`
	Name    string `json:"Name"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

/*
 * Slice of items to record item data
 */
var items = []Item{
	{ID: "0", Name: "Test Item 0", Desc: "Test Item Description", Content: "Wah!"},
	{ID: "1", Name: "Test Item 1", Desc: "Test Item Description", Content: "Guh!"},
	{ID: "2", Name: "Test Item 2", Desc: "Test Item Description", Content: "Peko!"},
}

/*
 * All items endpoint. Returns a json response of all items when hit.
 */
func allItems(c *gin.Context) {
	fmt.Println("Hit all items endpoint")
	c.IndentedJSON(http.StatusOK, items)
}

/*
 * Appends an item from JSON received in the request body.
 */
func postItem(c *gin.Context) {
	var newItem Item

	// Call BindJSON to bind the received JSON to newItem.
	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	items = append(items, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)
}

/*
 * Gets the item whose ID value matches the id given.
 */
func getItem(c *gin.Context) {
	id := c.Param("id")

	// Loops over the list of items to find an item with a matching ID value.
	for _, i := range items {
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

	// Loops over the list of items to find an item with a matching ID value.
	for index, i := range items {
		if i.ID == id {
			// Delete Item from slice
			items = append(items[:index], items[index+1:]...)

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
	param := strings.ToLower(c.Param("param"))
	newValue := c.Param("newValue")

	// Loops over the list of items to find an item with a matching ID value.
	for index, i := range items {
		if i.ID == id {
			// Patches based on param and newValue.
			switch param {
			case "name":
				i.Name = newValue
			case "desc":
				i.Desc = newValue
			case "content":
				i.Content = newValue
			default:
				c.IndentedJSON(http.StatusNotFound, gin.H{"message": "param not found"})
				return
			}

			items[index] = i
			c.IndentedJSON(http.StatusOK, i)
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

/*
 * Start Server.
 */
func startServer() {
	router := gin.Default()

	// Methods
	router.GET("/", homePage)
	router.GET("/items", allItems)
	router.GET("/items/:id", getItem)
	router.POST("/", homePagePOST)
	router.POST("/items", postItem)
	router.DELETE("items/:id", deleteItem)
	router.PATCH("items/:id/:param/:newValue", patchItem)

	router.Run()
}

/*
 * Main function.
 */
func main() {
	startServer()
}
