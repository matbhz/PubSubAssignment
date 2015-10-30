package Controllers

import (
	"github.com/gorilla/mux"
	"github.com/matbhz/CitrixAssignment/Utils"
	"net/http"
	"time"
	"github.com/matbhz/CitrixAssignment/Models"
)

const GET    = "GET"
const POST   = "POST"
const DELETE = "DELETE"

var Subscribers map[string]*Models.Subscriber

func StartSubscribers() {
	Subscribers = make(map[string]*Models.Subscriber)
}

func DefineRoutes() *mux.Router{
	r := mux.NewRouter()

	r.HandleFunc("/api/{topic}", Publish).Methods(POST).Headers("Content-Type", "application/json")

	r.HandleFunc("/api/{topic}/{subscriber}", Subscribe).Methods(POST)
	r.HandleFunc("/api/{topic}/{subscriber}", RemoveSubscriber).Methods(DELETE)
	r.HandleFunc("/api/{topic}/{subscriber}", Receive).Methods(GET)

	return r
}

func Publish(response http.ResponseWriter, request *http.Request) {
	publishedMessage := Utils.ParseMessage(request);

	if (publishedMessage == nil) {
		Utils.BadRequest("Message payload empty", response)
		return
	}

	publishedMessage.PublishedAt = time.Now().Format(time.RFC3339)

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
		Subscribers[name] = Models.NewSubscriber()
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
		return
	}

	message := Subscribers[name].Poll(topic)

	if (message == nil) {
		Utils.NotFound(response)
	} else {
		Utils.Ok(message, response)
	}
}