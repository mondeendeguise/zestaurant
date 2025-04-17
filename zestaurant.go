package main

import (
	"github.com/mondeendeguise/zestaurant/schema"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

//	"github.com/joho/godotenv"

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
	var unmarshaler json.Unmarshaler = conjson.NewUnmarshaler(&directory, transform.ConventionalKeys())
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

//	err := godotenv.Load(".env")
//	if err != nil {
//		// dont kill since this isnt really the end of the world i think
//		log.Println("WARNING:", err)
//	}

//	var apiUrl string = os.Getenv("API_URL")
//	if apiUrl == "" {
//		log.Println("WARNING: API_URL is not set")
//	}

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
			var directory schema.Directory = ReadMockDB(DB)
			var marshaler json.Marshaler = conjson.NewMarshaler(directory.Locations, transform.ConventionalKeys())
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
			var directory schema.Directory = ReadMockDB(DB)

			var locationName string = r.PathValue("location")
			var locationIndex int = -1
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

			var marshaler json.Marshaler = conjson.NewMarshaler(directory.Locations[locationIndex], transform.ConventionalKeys())
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
			var directory schema.Directory = ReadMockDB(DB)
			var locationName string = r.PathValue("location")
			var locationIndex int = -1
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

			var marshaler json.Marshaler = conjson.NewMarshaler(directory.Locations[locationIndex].Menu, transform.ConventionalKeys())
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

	http.HandleFunc("/locations/{location}/menu/{menuGroup}/", func(w http.ResponseWriter, r *http.Request) {
		switch(strings.ToUpper(r.Method)) {
		case "": fallthrough
		case "GET":
			var directory schema.Directory = ReadMockDB(DB)
			var locationName string = r.PathValue("location")
			var locationIndex int = -1
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

			var menuGroupName string = r.PathValue("menuGroup")
			var menuGroupIndex int = -1
			for i, v := range directory.Locations[locationIndex].Menu {
				if strings.EqualFold(v.Name, menuGroupName) {
					menuGroupIndex = i
					break
				}
			}

			if menuGroupIndex == -1 {
				http.Error(w, "404 Not Found", 404)
				return
			}

			var marshaler json.Marshaler = conjson.NewMarshaler(directory.Locations[locationIndex].Menu[menuGroupIndex], transform.ConventionalKeys())
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
