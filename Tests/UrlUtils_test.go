package Tests

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/matbhz/CitrixAssignment/Controllers"
	"log"
)

const TOPIC = "news"
const SUBSCRIBER = "Sub1"

func TestWhenSubscriberHasNoTopicsAndTriesToGetATopic_ShouldReturn404(t *testing.T){
	// Setup endpoints with real Controller routes for simplicity
	Controllers.StartSubscribers()
	r := Controllers.DefineRoutes()
	http.Handle("/", r)
	server := httptest.NewServer(r)
	defer server.Close()

	res, err := http.Get(server.URL+"/api/"+TOPIC+"/"+SUBSCRIBER)
	if err != nil {
		log.Fatal(err)
	}

	if (res.StatusCode != 404) {
		t.Error()
	}
}

func TestWhenSubscriberHasTopicsAndTriesToGetATopic_ShouldReturn200(t *testing.T){
	// Setup endpoints with real Controller routes for simplicity
	Controllers.StartSubscribers()
	r := Controllers.DefineRoutes()
	http.Handle("/", r)
	server := httptest.NewServer(r)
	defer server.Close()
}

func TestWhenSubscriberHasNoTopicsAndTriesToUnsubscribeFromATopic_ItShouldReturn204() {

}

func TestWhenSubscriberHasTopicsAndTriesToUnsubscribeFromATopic_ItShouldReturn204() {

}

func TestWhenSubscriberTriesSubscribeToATopic_ItShouldReturn201() {

}

func TestWhenPublisherPublishesAMessageInATopic_AllSubscribersShouldReceiveThatMessage() {

}
