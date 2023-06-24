package models

/*
 * Struct for a D&D character. Stored in a sqlite database.
 */
type Character struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"Name"`
	Race  string `json:"Race"`
	Class string `json:"Class"`
	Level uint8  `json:"Level"`
}
