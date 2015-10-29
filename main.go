package main

import (
	"net/http"
	"fmt"
	"log"
	"github.com/matbhz/CitrixAssignment/Models"
	"github.com/matbhz/CitrixAssignment/Controllers"
)

const GET    = "GET"
const POST   = "POST"
const DELETE = "DELETE"

func main() {

	r := Controllers.DefineRoutes()

	http.Handle("/", r)

	fmt.Println("tarting subscribers")

	Controllers.Subscribers = make(map[string]*Models.Subscriber)

	fmt.Println("Starting server locally on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}