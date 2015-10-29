package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"
	"github.com/matbhz/CitrixAssignment/Models"
	"github.com/matbhz/CitrixAssignment/Utils"
)

const GET    = "GET"
const POST   = "POST"
const DELETE = "DELETE"

var Subscribers map[string]*Models.Subscriber

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/{topic}", Publish).Methods(POST)

	r.HandleFunc("/api/{topic}/{subscriber}", Subscribe).Methods(POST)
	r.HandleFunc("/api/{topic}/{subscriber}", RemoveSubscriber).Methods(DELETE)
	r.HandleFunc("/api/{topic}/{subscriber}", Receive).Methods(GET)

	http.Handle("/", r)

	Subscribers = make(map[string]*Models.Subscriber)

	fmt.Println("Starting server locally on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Publish(response http.ResponseWriter, request *http.Request) {
	publishedMessage := Utils.ParseMessage(request);

	topic := Utils.GetTopic(request)

	for _, subscriber := range Subscribers {
		subscriberChannel := subscriber.Subscriptions[topic]
		if (subscriberChannel != nil){
			go func() {
				subscriberChannel <- publishedMessage
			}()
		}
	}

	Utils.NoResponse(response)
}

func Subscribe(response http.ResponseWriter, request *http.Request){
	topic, name := Utils.GetTopicAndSubscriber(request)

	if (Subscribers[name] == nil) {
		Subscribers[name] = Models.NewSubscriber(name)
	}

	if (Subscribers[name].HasSubscription(topic)) {
		Utils.BadRequest("Already subscribed to this topic", response)
	} else {
		Subscribers[name].Subscriptions[topic] = make(chan *Models.Message)
		Utils.Created(response)
	}
}

func RemoveSubscriber(response http.ResponseWriter, request *http.Request){
	topic, name := Utils.GetTopicAndSubscriber(request)

	if (Subscribers[name] == nil || !Subscribers[name].HasSubscription(topic)) {
		Utils.NotFound(response)
	} else {
		delete(Subscribers[name].Subscriptions, topic)
		Utils.NoResponse(response)
	}
}

func Receive(response http.ResponseWriter, request *http.Request){
	topic, name := Utils.GetTopicAndSubscriber(request)

	if (Subscribers[name] == nil || !Subscribers[name].HasSubscription(topic)) {
		Utils.NotFound(response)
	}

	message := Subscribers[name].Poll(topic)

	if (message == nil) {
		Utils.NotFound(response)
	} else {
		Utils.Ok(response)
	}
}