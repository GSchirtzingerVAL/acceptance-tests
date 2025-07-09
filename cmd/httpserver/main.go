package main

import (
	"log"
	"net/http"

	"github.com/quii/go-specs-greet/adapters/httpserver"
)

func main() {
	handler := httpserver.Handler()
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
