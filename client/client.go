package client

/*
	THIS PACKAGE IS ACTS AS AN ABSTRACT WAY OF INTERACTING WITH A YELLOW JACKET SERVER
*/

import (
	"net/http"

	"github.com/clavera2/yellow_jacket/utils"
)

type Client struct {
	http.Client
	serverURL string //the url of the server
}

var client Client

func initializeServerURL(url string) {
	//initializes client with the url of the server hosting messages
	client.serverURL = url
}

func MakeMessage(message interface{}) (*utils.Message, error) {
	m, err := utils.NewMessage(message)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func SendMessage(message *utils.Message) {
	//sends message to server, client MUST be connected to Yellow jacket server
}
