package Tests

import (
	"testing"
	"net/http"
	"fmt"
	"net/url"
	"net/http/httptest"
	"github.com/matbhz/CitrixAssignment/Controllers"
	"github.com/matbhz/CitrixAssignment/Utils"
)

func TestGetTopic_ShouldReturnTheTopic(t *testing.T){
	r := Controllers.DefineRoutes()
	http.Handle("/", r)

	server := httptest.NewServer(r)
	defer server.Close()

	x := &http.Request{
		URL: &url.URL{Path:"/api/A_COOL_TOPIC/something/els"},
		RequestURI:"/api/A_COOL_TOPIC/something/else"}

	result := Utils.GetTopic(x)

	fmt.Println(result)

}