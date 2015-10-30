# PubSub Assignment

## Endpoints:

# POST /api/:topic 

### Accepts: 
`Content-type` : `application/json`

### Action
Publishes a the a `message` to the `:topic` requested in the URI. 
Notifies *everyone* subscribed to that `topic`.

## Sample request payload:
```
#!json
{
   "message" : "I've got message for you!"
}
```

## Possible responses

* ### 204
Succesfully posted and notified subscribers with a message

* ### 404
If different, or empty, `Content-type` is requested.

# POST /api/:topic/:subscriber

### Action
Subscribes the `:subscriber` to a the given `:topic`

Messages published before the `:subscriber` has subscribed for that `:topic` will *not* be delivered.

## Possible responses

* ### 201
Subscribed to future messages in the `:topic`
	
# DELETE /api/:topic/:subscriber 

Unsubscribes the `:subscriber` from the given `:topic`
	
# GET /api/:topic/:subscriber

Tries to retrieve a message from the subscribed `:topic` for the given `:subscriber`.

Multiple calls are necessary to fully consume the subscription list. 

Returns `404` if all messages have been read or if the `:subscriber` is not subscribed to the given `:topic`.

## How to run from the source code:
On the root of the project, where `main.go` is located, run:

$ go run main.go

## Required  dependency
### Run:
$ go get github.com/gorilla/mux