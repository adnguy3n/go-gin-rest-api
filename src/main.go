package main

import (
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
var Items = []Item{
	{ID: "0", Name: "Test Item 0", Desc: "Test Item Description", Content: "Wah!"},
	{ID: "1", Name: "Test Item 1", Desc: "Test Item Description", Content: "Guh!"},
	{ID: "2", Name: "Test Item 2", Desc: "Test Item Description", Content: "Peko!"},
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
	router.PATCH("items/:id", patchItem)
	router.PUT("items/:id", putItem)

	router.Run()
}

/*
 * Main function.
 */
func main() {
	startServer()
}
