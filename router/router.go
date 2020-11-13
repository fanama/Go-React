package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Person ...
type Person struct {
	Name string `json:"name"`
}

// Form ...
func Form(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	var t Person
	if err != nil {
		panic(err)
	}
	// log.Println(string(body))
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	// fmt.Println("t :", t)
	fmt.Fprintln(w, t.Name)

}
