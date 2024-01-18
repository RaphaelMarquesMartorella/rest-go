package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/World", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	s := &http.Server{
		Addr:    ":3000",
		Handler: nil,
	}

	log.Fatal(s.ListenAndServe())
}
