package server

import (
	"net/http"
)

//THESE ARE THE HANDLERS THAT WILL HANDLE CLIENT REQUESTS

func homeHandler(res http.ResponseWriter, req *http.Request) {
	messages := mp.GetAllMessages()
}

func getAllMessagesHandler(res http.ResponseWriter, req *http.Request) {

}

func getMessageHandler(res http.ResponseWriter, req *http.Request) {

}

func addMessageHandler(res http.ResponseWriter, req *http.Request) {

}

func deleteMessageHandler(res http.ResponseWriter, req *http.Request) {

}
