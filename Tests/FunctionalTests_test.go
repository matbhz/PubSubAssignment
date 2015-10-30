package Tests

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/matbhz/CitrixAssignment/Controllers"
	"log"
	"bytes"
)

const NEWS_TOPIC = "news"
const SUBSCRIBER = "Sub1"

func TestWhenSubscriberHasNoTopicsAndTriesToGetMessages_ItShouldReturn404(t *testing.T){
	// Setup endpoints with real Controller routes for simplicity
	Controllers.StartSubscribers()
	r := Controllers.DefineRoutes()
	server := httptest.NewServer(r)
	defer server.Close()

	res, err := http.Get(server.URL+"/api/"+NEWS_TOPIC+"/"+SUBSCRIBER)
	if err != nil {
		log.Fatal(err)
	}

	if (res.StatusCode != 404) {
		t.Error()
	}
}

func TestWhenSubscriberHasSubscribedToTopicsWithoutMessagesAndTriesToGetMessagesFromTopic_ShouldReturn404(t *testing.T){
	// Setup endpoints with real Controller routes for simplicity
	Controllers.StartSubscribers()
	r := Controllers.DefineRoutes()
	server := httptest.NewServer(r)
	defer server.Close()

	_, _ = http.Post(server.URL+"/api/"+NEWS_TOPIC+"/"+SUBSCRIBER, "", nil)

	res, err := http.Get(server.URL+"/api/"+NEWS_TOPIC+"/"+SUBSCRIBER)
	if err != nil {
		log.Fatal(err)
	}

	if (res.StatusCode != 404) {
		t.Error("Was:", res.StatusCode)
	}
}

func TestWhenSubscriberHasSubscribedToTopicsWithMessagesAndTriesToGetMessagesFromTopic_ShouldReturn204(t *testing.T){
	// Setup endpoints with real Controller routes for simplicity
	Controllers.StartSubscribers()
	r := Controllers.DefineRoutes()
	http.Handle("/", r)
	server := httptest.NewServer(r)
	defer server.Close()

	_, _ = http.Post(server.URL+"/api/"+NEWS_TOPIC+"/"+SUBSCRIBER, "", nil)
	_, _ = http.Post(server.URL+"/api/"+NEWS_TOPIC+"/", "application/json", bytes.NewBuffer([]byte(`{ "message" : "This message is going to self-destruct in 5 seconds" }`)))

	res, err := http.Get(server.URL+"/api/"+NEWS_TOPIC+"/"+SUBSCRIBER)
	if err != nil {
		log.Fatal(err)
	}

	if (res.StatusCode != 404) {
		t.Error("Was:", res.StatusCode)
	}
}

func TestWhenSubscriberHasNoTopicsAndTriesToUnsubscribeFromATopic_ItShouldReturn404(t *testing.T) {
	// Setup endpoints with real Controller routes for simplicity
	Controllers.StartSubscribers()
	r := Controllers.DefineRoutes()
	server := httptest.NewServer(r)
	defer server.Close()

	req, err := http.NewRequest("DELETE", server.URL+"/api/"+NEWS_TOPIC+"/"+SUBSCRIBER, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if (res.StatusCode != 404) {
		t.Error("Was:", res.StatusCode)
	}
}

func TestWhenSubscriberHasSubscribedToATopicAndTriesToUnsubscribeFromIt_ItShouldReturn204(t *testing.T) {
	// Setup endpoints with real Controller routes for simplicity
	Controllers.StartSubscribers()
	r := Controllers.DefineRoutes()
	server := httptest.NewServer(r)
	defer server.Close()

	_, _ = http.Post(server.URL+"/api/"+NEWS_TOPIC+"/"+SUBSCRIBER, "", nil)


	req, err := http.NewRequest("DELETE", server.URL+"/api/"+NEWS_TOPIC+"/"+SUBSCRIBER, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if (res.StatusCode != 204) {
		t.Error("Was:", res.StatusCode)
	}
}

func TestWhenSubscriberSubscribesToATopic_ItShouldReturn201(t *testing.T) {
	// Setup endpoints with real Controller routes for simplicity
	Controllers.StartSubscribers()
	r := Controllers.DefineRoutes()
	server := httptest.NewServer(r)
	defer server.Close()

	res, err := http.Post(server.URL+"/api/"+NEWS_TOPIC+"/"+SUBSCRIBER, "", nil)
	if err != nil {
		log.Fatal(err)
	}

	if (res.StatusCode != 201) {
		t.Error("Was:", res.StatusCode)
	}
}

func TestWhenSubscriberSubscribesToATopicMoreThanOnce_ItShouldReturn400(t *testing.T) {
	// Setup endpoints with real Controller routes for simplicity
	Controllers.StartSubscribers()
	r := Controllers.DefineRoutes()
	server := httptest.NewServer(r)
	defer server.Close()

	_, _ = http.Post(server.URL+"/api/"+NEWS_TOPIC+"/"+SUBSCRIBER, "", nil)
	res, err := http.Post(server.URL+"/api/"+NEWS_TOPIC+"/"+SUBSCRIBER, "", nil)
	if err != nil {
		log.Fatal(err)
	}

	if (res.StatusCode != 400) {
		t.Error("Was:", res.StatusCode)
	}
}

func TestWhenPublisherPublishesToATopic_ItShouldReturn204(t *testing.T) {
	// Setup endpoints with real Controller routes for simplicity
	Controllers.StartSubscribers()
	r := Controllers.DefineRoutes()
	server := httptest.NewServer(r)
	defer server.Close()

	res, err := http.Post(server.URL+"/api/"+NEWS_TOPIC, "application/json", bytes.NewBuffer([]byte(`{ "message" : "You like them news, eh?" }`)))
	if err != nil {
		log.Fatal(err)
	}

	if (res.StatusCode != 204) {
		t.Error("Was:", res.StatusCode)
	}
}