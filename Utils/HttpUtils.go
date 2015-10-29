package Utils
import "net/http"

func Ok(response http.ResponseWriter) {
	response.WriteHeader(200)
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