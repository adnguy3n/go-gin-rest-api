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
	Name    string `json:"Name"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

/*
 * Slice of items to record item data
 */
var items = []item{
	{Name: "Test Item", Desc: "Test Item Description", Content: "Wah!"},
}

/*
 * All items endpoint. Returns a json response of all items when hit.
 */
func allItems(c *gin.Context) {
	fmt.Println("Hit all items endpoint")
	c.IndentedJSON(http.StatusOK, items)
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
 * Main function
 */
func main() {
	router := gin.Default()
	router.GET("/", homePage)
	router.GET("/items", allItems)

	router.Run()
}
