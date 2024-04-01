package main

import (
	"encoding/json"
	"github.com/go-faker/faker/v4"
	"net/http"
)

type Person struct {
	FirstName string `faker:"first_name"`
	LastName  string `faker:"last_name"`
}

func personsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/persons" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	p := Person{}
	if err := faker.FakeData(&p); err != nil {
		http.Error(w, err.Error(), http.StatusOK)
		return
	}
	json.NewEncoder(w).Encode(p)
}

func main() {
	http.HandleFunc("/person", personsHandler)
	http.ListenAndServe(":8080", nil)
}
