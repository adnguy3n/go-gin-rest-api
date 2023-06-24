package models

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
