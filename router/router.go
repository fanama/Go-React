package router

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Handler ...
func Handler(w http.ResponseWriter, r *http.Request) {

	type Person struct {
		Name     string
		LastName string
		Age      uint8
	}

	person := Person{Name: "Shashank", LastName: "Tiwari", Age: 30}

	jsonResponse, jsonError := json.Marshal(person)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// Form ...
func Form(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)

}
