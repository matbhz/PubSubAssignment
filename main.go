package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"
)

const GET    = "GET"
const POST   = "POST"
const DELETE = "DELETE"

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/{topic}", Publish).Methods(POST)

	r.HandleFunc("/api/{topic}/{subscriber}", Subscribe).Methods(POST)
	r.HandleFunc("/api/{topic}/{subscriber}", RemoveSubscriber).Methods(DELETE)
	r.HandleFunc("/api/{topic}/{subscriber}", Receive).Methods(GET)

	http.Handle("/", r)

	fmt.Print("Starting server locally on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}


func Publish(http.ResponseWriter, *http.Request){
	fmt.Print("POSTed Publish")
}

func Subscribe(http.ResponseWriter, *http.Request){
	fmt.Print("POSTed Subscribe")
}

func RemoveSubscriber(http.ResponseWriter, *http.Request){
	fmt.Print("DELETEd RemoveSubscriber")
}

func Receive(http.ResponseWriter, *http.Request){
	fmt.Print("GET Receive")
}