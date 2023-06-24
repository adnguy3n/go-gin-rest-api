package main

import (
	"go-gin-rest-api/src/controllers"

	"github.com/gin-gonic/gin"
)

/*
 * Start Server.
 */
func startServer() {
	router := gin.Default()

	// Methods
	router.GET("/", controllers.HomePage)
	router.GET("/items", controllers.AllItems)
	router.GET("/items/:id", controllers.GetItem)
	router.POST("/", controllers.HomePagePOST)
	router.POST("/items", controllers.PostItem)
	router.DELETE("items/:id", controllers.DeleteItem)
	router.PATCH("items/:id", controllers.PatchItem)
	router.PUT("items/:id", controllers.PutItem)

	router.Run()
}

/*
 * Main function.
 */
func main() {
	startServer()
}
