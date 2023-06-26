package databases

import (
	"go-gin-rest-api/src/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

/*
 * Defining the database. Used across the entire app to communicate with the Database.
 */
var CharacterDB *gorm.DB

/*
 * Connects to the Database or creates it if it doesn't exist.
 * Currently set to test.db which can be found in the database folder.
 */
func ConnectDatabase() {

	database, err := gorm.Open(sqlite.Open("database/characters.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&models.Character{})
	if err != nil {
		return
	}

	CharacterDB = database
}
