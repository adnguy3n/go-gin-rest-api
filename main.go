package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/*
 * Item data structure.
 */
type Item struct {
	Name    string `json:"Name"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

/*
 * Array of Items.
 */
type Items []Item

/*
 * All items endpoint. Returns a json response of all items when hit.
 */
func allItems(w http.ResponseWriter, r *http.Request) {
	items := Items{
		Item{Name: "Test Item", Desc: "Test Item Description", Content: "Wah!"},
	}

	fmt.Println("Hit all items endpoint")
	json.NewEncoder(w).Encode(items)
}

/*
 * HomePage endpoint. Prints out a message when hit.
 */
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, this is Home")
}

/*
 * Main function
 */
func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/items", allItems)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
