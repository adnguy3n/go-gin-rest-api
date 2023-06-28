package main

import (
	"go-gin-rest-api/src/controllers"
	"go-gin-rest-api/src/databases"
	"go-gin-rest-api/src/middlewares"

	"github.com/gin-gonic/gin"
)

/*
 * Start Server.
 */
func startServer() {
	// Initialize Databases.
	databases.ConnectCharacterDatabase()
	databases.ConnectUserDatabase("root:rootpassword@tcp(userdb:3306)/userdb?parseTime=true")
	databases.MigrateUserDatabase()

	router := initRouter()
	router.Run()
}

/*
 * Initialize router.
 */
func initRouter() *gin.Engine {
	router := gin.Default()

	// Methods for HomePage endpoints.
	router.GET("/", controllers.HomePage)
	router.POST("/", controllers.HomePagePOST)

	// Router Groups
	characters(router)
	users(router)

	return router
}

/*
 * Router Group for Characters.
 * Uses GORM and SQLite.
 * Create, Edit, and Delete D&D Characters in database.
 */
func characters(router *gin.Engine) {
	// Paths with this will take an id parameter.
	const id = "/:id"

	characters := router.Group("/characters")
	{
		characters.GET("/", controllers.AllCharacters)
		characters.GET(id, controllers.GetCharacter)

		// Requires JWT authenthication to create, update, or delete characters.
		secured := characters.Group("/").Use(middlewares.Authenthicate())
		{
			secured.POST("/", controllers.PostCharacter)
			secured.DELETE(id, controllers.DeleteCharacter)
			secured.PATCH(id, controllers.PatchCharacter)
			secured.PUT(id, controllers.PutCharacter)
		}
	}
}

/*
 * Router Group for Users.
 * Uses GORM and MySql.
 * Register users and generate tokens.
 */
func users(router *gin.Engine) {
	// Paths with this will take a username parameter.
	const username = "/username"

	users := router.Group("/users")
	{
		users.GET(username, controllers.GetUser)
		users.POST("/", controllers.RegisterUser)
		users.POST("/token", controllers.GenerateToken)

		// Requires JWT authenthication to update or delete characters.
		secured := users.Group("/").Use(middlewares.Authenthicate())
		{
			secured.PATCH(username, controllers.UpdateUser)
			secured.DELETE(username, controllers.DeleteUser)
		}
	}
}

/*
 * Main function.
 */
func main() {
	startServer()
}
