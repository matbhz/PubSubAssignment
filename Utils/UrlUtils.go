package Utils
import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/matbhz/CitrixAssignment/Models"
)

func GetTopic(request *http.Request) (string) {
	uriVariables := mux.Vars(request)
	return uriVariables["topic"]
}

func GetTopicAndSubscriber(request *http.Request) (string, string) {
	uriVariables := mux.Vars(request)
	return uriVariables["topic"], uriVariables["subscriber"]
}

func ParseMessage(request *http.Request) *Models.Message {
	decoder := json.NewDecoder(request.Body)
	var publishedMessage *Models.Message
	err := decoder.Decode(&publishedMessage)
	if err != nil {
		panic(err)
	}
	return publishedMessage
}