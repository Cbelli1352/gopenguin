package main

import (
	"log"
	"net/http"

	"gopenguin/route"
)

func main() {
	route.Load()

	log.Print("Starting server")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
