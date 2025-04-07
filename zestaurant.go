package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Directory struct {
	Locations Locations
}

type Location struct {
	Name string
	Menu Menu
}
type Locations []Location

type Menu struct {
	Appetizers MenuItems
	Soup MenuItems
	Entrees MenuItems
	Desserts MenuItems
	Beverages Strings
}

type Strings []string

type MenuItem struct {
	Name string
	Price float64
	Description string
	InStock bool
}
type MenuItems []MenuItem

const DB = "db.json"

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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.MarshalIndent(directory, "", "\t")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, string(b))
	})

	// TODO: handle requests to paths with trailing `/`
	http.HandleFunc("/locations", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.MarshalIndent(directory.Locations, "", "\t")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, string(b))
	})

	http.HandleFunc("/locations/{location}", func(w http.ResponseWriter, r *http.Request) {
		location := r.PathValue("location")

		locationIndex := -1
		for i, v := range directory.Locations {
			if strings.EqualFold(v.Name, location) {
				locationIndex = i
			}
		}

		if locationIndex == -1 {
			fmt.Fprintf(w, "not found")
			return
		}

		b, err := json.MarshalIndent(directory.Locations[locationIndex], "", "\t")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, string(b))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
