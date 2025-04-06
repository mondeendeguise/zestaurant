package main

import (
	"encoding/json"
	"fmt"
	// "html"
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

func (directory Directory) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b, err := json.MarshalIndent(directory, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, string(b))
}

func (location Location) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func (menu Menu) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func (menuItem MenuItem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(menuItem)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "%q", b)
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

	http.Handle("/", directory)
	http.Handle("/directory", directory)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
