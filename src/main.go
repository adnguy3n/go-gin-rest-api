package main

import (
	"go-gin-rest-api/src/controllers"
	"go-gin-rest-api/src/databases"

	"github.com/gin-gonic/gin"
)

/*
 * Start Server.
 */
func startServer() {
	router := gin.Default()
	databases.ConnectCharacterDatabase()
	databases.ConnectUserDatabase("root:rootpassword@tcp(userdb:3306)/userdb?parseTime=true")
	databases.MigrateUserDatabase()

	// Methods for HomePage endpoints.
	router.GET("/", controllers.HomePage)
	router.POST("/", controllers.HomePagePOST)

	// Methods for D&D Characters. Uses GORM and SQLite.
	router.GET("/characters", controllers.AllCharacters)
	router.GET("/characters/:id", controllers.GetCharacter)
	router.POST("/characters", controllers.PostCharacter)
	router.DELETE("/characters/:id", controllers.DeleteCharacter)
	router.PATCH("/characters/:id", controllers.PatchCharacter)
	router.PUT("/characters/:id", controllers.PutCharacter)

	// Methods for Users.
	router.POST("/users", controllers.RegisterUser)
	//router.GET("/users", controllers.AllUsers)

	// Methods for Item. Uses a slice.
	router.GET("/items", controllers.AllItems)
	router.GET("/items/:id", controllers.GetItem)
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
