package main

import (
	"github.com/mondeendeguise/zestaurant/schema"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"github.com/Rican7/conjson"
	"github.com/Rican7/conjson/transform"
)

// TODO: simulate external API/DB
func ReadMockDB(file string) schema.Directory {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	directory := schema.Directory{}
	unmarshaler := conjson.NewUnmarshaler(&directory, transform.ConventionalKeys())
	err = json.Unmarshal(data, unmarshaler)
	if err != nil {
		log.Fatal(err)
	}

	return directory
}

const HTTP_STATUS_NOT_IMPLEMENTED = 501
const HTTP_STATUS_METHOD_NOT_ALLOWED = 405

const DB = "db.json"

func main() {
	log.SetOutput(os.Stderr)
	log.SetFlags(log.Llongfile)

	err := godotenv.Load(".env")
	if err != nil {
		// dont kill since this isnt really the end of the world i think
		log.Println("WARNING:", err)
	}

	apiUrl := os.Getenv("API_URL")
	if apiUrl == "" {
		log.Println("WARNING: API_URL is not set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch strings.ToUpper(r.Method) {
		case "": fallthrough
		case "GET":
			fmt.Fprintf(w, "{}")

		default:
			http.Error(w, "Method Not Allowed", HTTP_STATUS_METHOD_NOT_ALLOWED)
			return
		}
	})

	// TODO: handle requests to paths with trailing `/`
	// TODO: make paths case insensitive

	http.HandleFunc("/locations/", func(w http.ResponseWriter, r *http.Request) {
		switch(strings.ToUpper(r.Method)) {
		case "": fallthrough
		case "GET":
			directory := ReadMockDB(DB)
			marshaler := conjson.NewMarshaler(directory.Locations, transform.ConventionalKeys())
			b, err := json.Marshal(marshaler)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Fprintf(w, "%s", string(b))

		case "POST":
			http.Error(w, "POST Not Implemented", HTTP_STATUS_NOT_IMPLEMENTED)
			return

		case "PUT":
			http.Error(w, "PUT Not Implemented", HTTP_STATUS_NOT_IMPLEMENTED)
			return

		case "PATCH":
			http.Error(w, "PATCH Not Implemented", HTTP_STATUS_NOT_IMPLEMENTED)
			return

		case "DELETE":
			http.Error(w, "DELETE Not Implemented", HTTP_STATUS_NOT_IMPLEMENTED)
			return

		default:
			http.Error(w, "Method Not Allowed", HTTP_STATUS_METHOD_NOT_ALLOWED)
			return
		}
	})

	http.HandleFunc("/locations/{location}/", func(w http.ResponseWriter, r *http.Request) {
		switch(strings.ToUpper(r.Method)) {
		case "": fallthrough
		case "GET":
			directory := ReadMockDB(DB)

			locationName := r.PathValue("location")
			locationIndex := -1
			for i, v := range directory.Locations {
				if strings.EqualFold(v.Name, locationName) {
					locationIndex = i
					break
				}
			}

			if locationIndex == -1 {
				http.Error(w, "404 Not Found", 404)
				return
			}

			marshaler := conjson.NewMarshaler(directory.Locations[locationIndex], transform.ConventionalKeys())
			b, err := json.Marshal(marshaler)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Fprintf(w, "%s", string(b))

		case "POST":
			http.Error(w, "POST Not Implemented", HTTP_STATUS_NOT_IMPLEMENTED)
			return

		case "PUT":
			http.Error(w, "PUT Not Implemented", HTTP_STATUS_NOT_IMPLEMENTED)
			return

		case "PATCH":
			http.Error(w, "PATCH Not Implemented", HTTP_STATUS_NOT_IMPLEMENTED)
			return

		case "DELETE":
			http.Error(w, "DELETE Not Implemented", HTTP_STATUS_NOT_IMPLEMENTED)
			return

		default:
			http.Error(w, "Method Not Allowed", HTTP_STATUS_METHOD_NOT_ALLOWED)
			return
		}
	})

	http.HandleFunc("/locations/{location}/menu/", func(w http.ResponseWriter, r *http.Request) {
		switch(strings.ToUpper(r.Method)) {
		case "":
		case "GET":
			directory := ReadMockDB(DB)
			locationName := r.PathValue("location")
			locationIndex := -1
			for i, v := range directory.Locations {
				if strings.EqualFold(v.Name, locationName) {
					locationIndex = i
					break
				}
			}

			if locationIndex == -1 {
				http.Error(w, "404 Not Found", 404)
				return
			}

			marshaler := conjson.NewMarshaler(directory.Locations[locationIndex].Menu, transform.ConventionalKeys())
			b, err := json.Marshal(marshaler)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Fprintf(w, "%s", string(b))

		case "POST":
			http.Error(w, "POST Not Implemented", HTTP_STATUS_NOT_IMPLEMENTED)
			return

		case "PUT":
			http.Error(w, "PUT Not Implemented", HTTP_STATUS_NOT_IMPLEMENTED)
			return

		case "PATCH":
			http.Error(w, "PATCH Not Implemented", HTTP_STATUS_NOT_IMPLEMENTED)
			return

		case "DELETE":
			http.Error(w, "DELETE Not Implemented", HTTP_STATUS_NOT_IMPLEMENTED)
			return

		default:
			http.Error(w, "Method Not Allowed", HTTP_STATUS_METHOD_NOT_ALLOWED)
			return
		}
	})

	http.HandleFunc("/locations/{location}/menu/{subMenu}/", func(w http.ResponseWriter, r *http.Request) {
		switch(strings.ToUpper(r.Method)) {
		case "": fallthrough
		case "GET":
			directory := ReadMockDB(DB)
			locationName := r.PathValue("location")
			locationIndex := -1
			for i, v := range directory.Locations {
				if strings.EqualFold(v.Name, locationName) {
					locationIndex = i
					break
				}
			}

			if locationIndex == -1 {
				http.Error(w, "404 Not Found", 404)
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
				http.Error(w, "404 Not Found", 404)
				return
			}

			marshaler := conjson.NewMarshaler(directory.Locations[locationIndex].Menu[subMenuIndex], transform.ConventionalKeys())
			b, err := json.Marshal(marshaler)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Fprintf(w, "%s", string(b))

		case "POST":
			http.Error(w, "POST Not Implemented", HTTP_STATUS_NOT_IMPLEMENTED)
			return

		case "PUT":
			http.Error(w, "PUT Not Implemented", HTTP_STATUS_NOT_IMPLEMENTED)
			return

		case "PATCH":
			http.Error(w, "PATCH Not Implemented", HTTP_STATUS_NOT_IMPLEMENTED)
			return

		case "DELETE":
			http.Error(w, "DELETE Not Implemented", HTTP_STATUS_NOT_IMPLEMENTED)
			return

		default:
			http.Error(w, "Method Not Allowed", HTTP_STATUS_METHOD_NOT_ALLOWED)
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
