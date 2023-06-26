package controllers

import (
	"go-gin-rest-api/src/databases"
	"go-gin-rest-api/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
 * Struct for a D&D character.
 * A unique ID is not needed as it will be automatically generated.
 */
type inputCharacter struct {
	Name  string `json:"Name"`
	Race  string `json:"Race"`
	Class string `json:"Class"`
	Level uint8  `json:"level"`
}

/*
 * All items endpoint. Returns a json response of all characters when hit.
 */
func AllCharacters(c *gin.Context) {
	var characters []models.Character
	databases.CharacterDB.Find(&characters)
	c.IndentedJSON(http.StatusOK, characters)
}

/*
 * Gets the character whose ID value matches the id given.
 */
func GetCharacter(c *gin.Context) {
	var character models.Character

	// Finds the character based on their unique ID.
	// Gives an error if no character with that ID exists.
	err := databases.CharacterDB.Where("id = ?", c.Param("id")).First(&character).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Character not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, character)
}

/*
 * Appends an Character from JSON received in the request body.
 */
func PostCharacter(c *gin.Context) {
	var newCharacter inputCharacter

	// Call BindJSON to bind the received JSON to newCharacter.
	if err := c.BindJSON(&newCharacter); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Stores the new character's Name, Race, Class, and Level in character.
	character := models.Character{
		Name:  &newCharacter.Name,
		Race:  &newCharacter.Race,
		Class: &newCharacter.Race,
		Level: &newCharacter.Level}

	// The max level for D&D character is 20.
	if *character.Level > 20 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A character cannot go past 20th level."})
		return
	}

	databases.CharacterDB.Create(&character)

	c.IndentedJSON(http.StatusCreated, character)
}

/*
 * Deletes the character whose ID value matches the id given.
 */
func DeleteCharacter(c *gin.Context) {
	var character models.Character

	// Finds the character based on their unique ID.
	// Gives an error if no character with that ID exists.
	err := databases.CharacterDB.Where("id = ?", c.Param("id")).First(&character).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Character not found."})
		return
	}

	databases.CharacterDB.Delete(&character)
	c.JSON(http.StatusNotFound, gin.H{"Message": "Character deleted."})
}

/*
 * Patches the character whose ID value matches the id given.
 */
func PatchCharacter(c *gin.Context) {
	var character models.Character
	var updatedCharacter inputCharacter

	// Finds the character based on their unique ID.
	// Gives an error if no character with that ID exists.
	err := databases.CharacterDB.Where("id = ?", c.Param("id")).First(&character).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Character not found."})
		return
	}

	// Call BindJSON to bind the received JSON to updatedCharacter.
	if err := c.BindJSON(&updatedCharacter); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "input not found"})
		return
	}

	// The max level for D&D character is 20.
	if updatedCharacter.Level > 20 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A character cannot go past 20th level."})
		return
	}

	databases.CharacterDB.Model(&character).Updates(updatedCharacter)
	c.IndentedJSON(http.StatusOK, character)
}

/*
 * PUT the item whose ID value matches the id given.
 */
func PutCharacter(c *gin.Context) {
	var character models.Character
	var updatedCharacter models.Character

	// Finds the character based on their unique ID.
	// Gives an error if no character with that ID exists.
	err := databases.CharacterDB.Where("id = ?", c.Param("id")).First(&character).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Character not found."})
		return
	}

	// Call BindJSON to bind the received JSON to updatedCharacter.
	if err := c.BindJSON(&updatedCharacter); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "input not found"})
		return
	}

	if updatedCharacter.ID == nil || updatedCharacter.Name == nil || updatedCharacter.Race == nil ||
		updatedCharacter.Class == nil || updatedCharacter.Level == nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "PUT requests requires all fields to be filled."})
		return
	}

	// The max level for D&D character is 20.
	if *updatedCharacter.Level > 20 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A character cannot go past 20th level."})
		return
	}

	databases.CharacterDB.Model(&character).Updates(updatedCharacter)
	c.IndentedJSON(http.StatusOK, character)
}
