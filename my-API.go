package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"./auth"
	"./router"
)

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")

		println(authHeader)

		if len(authHeader) != 2 {
			println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]

			println(jwtToken)

		}

	})
}

func main() {

	fileServer := http.FileServer(http.Dir("./public"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", router.Form)
	http.HandleFunc("/auth/new", auth.New)
	http.Handle("/middleware", middleware(http.HandlerFunc(router.Handler)))

	http.Handle("/sql/protected", middleware(http.HandlerFunc(router.Handler)))

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
