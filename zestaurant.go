package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// TODO: simulate external API/DB
//       `db.json` is read into memory once at startup
//       make a `fetch` abstraction for requesting arbitrary data methinks
const DB = "db.json"

type Directory struct {
	Locations Locations
}

type Location struct {
	Name string
	Menu SubMenus
}
type Locations []Location

type SubMenu struct {
	Name string
	Items MenuItems
}
type SubMenus []SubMenu

type MenuItem struct {
	Name string
	Price float64
	Description string
	InStock bool
}
type MenuItems []MenuItem

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Llongfile)

	data, err := os.ReadFile("db.json")
	if err != nil {
		log.Fatal(err)
	}

	directory := Directory{}

	err = json.Unmarshal(data, &directory)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: handle POST, PUT, PATCH, DELETE
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(directory)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, string(b))
	})

	// TODO: handle requests to paths with trailing `/`
	// TODO: make paths case insensitive
	http.HandleFunc("/locations", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(directory.Locations)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, string(b))
	})

	http.HandleFunc("/locations/{location}", func(w http.ResponseWriter, r *http.Request) {

		locationName := r.PathValue("location")
		locationIndex := -1
		for i, v := range directory.Locations {
			if strings.EqualFold(v.Name, locationName) {
				locationIndex = i
				break
			}
		}

		if locationIndex == -1 {
			http.Error(w, "not found", 404)
			return
		}

		b, err := json.Marshal(directory.Locations[locationIndex])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, string(b))
	})

	http.HandleFunc("/locations/{location}/menu", func(w http.ResponseWriter, r *http.Request) {
		locationName := r.PathValue("location")
		locationIndex := -1
		for i, v := range directory.Locations {
			if strings.EqualFold(v.Name, locationName) {
				locationIndex = i
				break
			}
		}

		if locationIndex == -1 {
			http.Error(w, "not found", 404)
			return
		}

		b, err := json.Marshal(directory.Locations[locationIndex].Menu)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, string(b))
	})

	http.HandleFunc("/locations/{location}/menu/{subMenu}", func(w http.ResponseWriter, r *http.Request) {
		locationName := r.PathValue("location")
		locationIndex := -1
		for i, v := range directory.Locations {
			if strings.EqualFold(v.Name, locationName) {
				locationIndex = i
				break
			}
		}

		if locationIndex == -1 {
			http.Error(w, "not found", 404)
			return
		}

		subMenuName := r.PathValue("subMenu")
		subMenuIndex := -1
		for i, v := range directory.Locations[locationIndex].Menu {
			if strings.EqualFold(v.Name, subMenuName) {
				subMenuIndex = i
				break
			}
		}

		if subMenuIndex == -1 {
			http.Error(w, "not found", 404)
			return
		}

		b, err := json.Marshal(directory.Locations[locationIndex].Menu[subMenuIndex])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, string(b))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
