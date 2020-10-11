package main

import (
	"fmt"
	"log"
	"net/http"

	"./router"
)

func main() {

	fileServer := http.FileServer(http.Dir("./public"))

	http.Handle("/", fileServer)
	http.HandleFunc("/test", router.Handler)
	http.HandleFunc("/form", router.Form)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
