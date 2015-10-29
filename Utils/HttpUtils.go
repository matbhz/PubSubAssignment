package Utils
import (
	"net/http"
	"encoding/json"
)

func Ok(message interface{}, response http.ResponseWriter) {
	response.WriteHeader(200)
	serializedMessage, _ := json.Marshal(message)  // TODO: Unlikely, but treat error better
	response.Write([]byte(serializedMessage))
}

func Created(response http.ResponseWriter) {
	response.WriteHeader(201)
}

func NoResponse(response http.ResponseWriter) {
	response.WriteHeader(204)
}

func BadRequest(message string, response http.ResponseWriter) {
	response.WriteHeader(400)
	response.Write([]byte(message))
}

func NotFound(response http.ResponseWriter) {
	response.WriteHeader(404)
}