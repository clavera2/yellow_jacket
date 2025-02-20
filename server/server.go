package server

import (
	"net/http"
)

var server = http.NewServeMux()

var mP = NewMessagePool() //will hold messages owned by client

func main() {
	//the entry point to the server application
}

func initializeServer() {
	//add all routes and handlers to the server....get it ready for client requests
}
