package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiRoute struct {
	Endpoint string
	Verb     string
	Handler  func(http.ResponseWriter, *http.Request)
}

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

func initData() []Person {
	var people []Person

	people = append(people, Person{
		ID:        "1",
		Firstname: "John",
		Lastname:  "Doe",
		Address: &Address{
			City:  "City X",
			State: "State X",
		},
	})

	people = append(people, Person{
		ID:        "2",
		Firstname: "Koko",
		Lastname:  "Doe",
		Address: &Address{
			City:  "City Z",
			State: "State Y",
		},
	})

	people = append(people, Person{
		ID:        "3",
		Firstname: "Francis",
		Lastname:  "Sunday",
	})

	return people
}

func initRoutes(database []Person) []ApiRoute {
	var routes []ApiRoute

	routes = append(routes, ApiRoute{
		Endpoint: "/people",
		Verb:     "GET",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(database)
			return
		},
	})

	routes = append(routes, ApiRoute{
		Endpoint: "/people/{id}",
		Verb:     "GET",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			var requestedPerson Person

			params := mux.Vars(r)
			for _, person := range database {
				if person.ID == params["id"] {
					requestedPerson = person
					break
				}
			}

			json.NewEncoder(w).Encode(requestedPerson)
			return
		},
	})

	return routes
}

func registerRoutes(router *mux.Router, routes []ApiRoute) {
	for _, route := range routes {
		router.HandleFunc(route.Endpoint, route.Handler).Methods(route.Verb)
	}
}

func main() {
	database := initData()

	router := mux.NewRouter()
	routes := initRoutes(database)

	registerRoutes(router, routes)

	log.Fatal(http.ListenAndServe(":3030", router))
}
