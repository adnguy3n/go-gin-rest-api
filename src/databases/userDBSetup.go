package databases

import (
	"go-gin-rest-api/src/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Defines the user database. Used to commmunicate with the database.
var UserDB *gorm.DB
var dbError error

/*
 * Connection function takes in the MySQL connection string and will try to
 * connect to the database using GORM.
 */
func ConnectUserDatabase(connectionString string) {
	UserDB, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to database")
	}
	log.Println("Connected to Database.")
}

/*
 * Migrate function ensures that there is a users table in the database.
 * If a users table is not in the database, GORM will automatically create
 * a new users table.
 */
func MigrateUserDatabase() {
	UserDB.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed.")
}
