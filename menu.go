package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

type Directory struct {
	Locations []Location
}

type Location struct {
	Name string
	Menu Menu
}

type Menu struct {
	Appetizers []MenuItem
	Soup []MenuItem
	Entrees []MenuItem
	Desserts []MenuItem
	Beverages []string
}

type MenuItem struct {
	Name string
	Price float64
	Description string
	InStock bool
}

const DB = "db.json"

func (d Directory) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	data, err := os.ReadFile("db.json")
	if err != nil {
		log.Fatal(err)
	}

	directory := Directory{}

	err = json.Unmarshal(data, &directory)
	if err != nil {
		log.Fatal(err)
	}

	for location := range directory.locations {
		
	}

	http.Handle("/directory", directory)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
