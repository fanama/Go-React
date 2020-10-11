package router

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type person struct {
	Name     string
	LastName string
	Age      uint8
}

// Handler ...
func Handler(w http.ResponseWriter, r *http.Request) {
	person := person{Name: "Shashank", LastName: "Tiwari", Age: 30}

	jsonResponse, jsonError := json.Marshal(person)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	fmt.Println(string(jsonResponse))

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
