package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
 * Item data structure.
 */
type item struct {
	ID      string `json:"id"`
	Name    string `json:"Name"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

/*
 * Slice of items to record item data
 */
var items = []item{
	{ID: "0", Name: "Test Item", Desc: "Test Item Description", Content: "Wah!"},
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
	var newItem item

	// Call BindJSON to bind the received JSON to newItem.
	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	items = append(items, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)
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
	router.GET("/", homePage)
	router.GET("/items", allItems)
	router.POST("/", homePagePOST)
	router.POST("/items", postItem)
	router.Run()
}

/*
 * Main function.
 */
func main() {
	startServer()
}
