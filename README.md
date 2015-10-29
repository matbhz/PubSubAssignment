# PubSub Assignment

## Endpoints:

# POST /api/:topic

## Sample request payload:
```
#!json
{
   "message" : "I've got message for you!"
}
```

Publishes a the a `message` to the `:topic` requested in the URI. 
Notifies *everyone* subscribed to that `topic`.
 
# POST /api/:topic/:subscriber

Subscribes the `:subscriber` to a the given `:topic`

Messages published before the `:subscriber` has subscribed for that `:topic` will *not* be delivered.
	
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