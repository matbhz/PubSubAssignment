package main

import (
	"net/http"
	"fmt"
	"log"
	"github.com/matbhz/CitrixAssignment/Controllers"
)

func main() {

	r := Controllers.DefineRoutes()

	http.Handle("/", r)
	Controllers.StartSubscribers()

	fmt.Println("Starting server locally on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}