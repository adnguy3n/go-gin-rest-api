package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage( w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, this is Home")
}

func main() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}