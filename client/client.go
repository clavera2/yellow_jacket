package client

/*
	THIS PACKAGE ACTS AS AN ABSTRACT WAY OF INTERACTING WITH A YELLOW JACKET SERVER
	IT IS INITIALLY USED AS A CLI CLIENT APPLICATION
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/clavera2/yellow_jacket/utils"
)

type Client struct {
	http.Client
	serverURL string // the URL of the server
}

var client Client

func InitializeServerURL(url string) {

	client.serverURL = url
}

func MakeMessage(content interface{}) (*utils.Message, error) {
	m, err := utils.NewMessage(content)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func SendMessage(message *utils.Message) error {
	if client.serverURL == "" {
		return fmt.Errorf("server URL is not initialized")
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	resp, err := client.Post(client.serverURL+"/add", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to send message: status %s", resp.Status)
	}

	fmt.Println("✅ Message sent successfully.")
	return nil
}

func GetMessageByID(id string) error {
	resp, err := client.Get(client.serverURL + "/get?id=" + id)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server responded with: %s", resp.Status)
	}

	var msg utils.Message
	if err := json.NewDecoder(resp.Body).Decode(&msg); err != nil {
		return err
	}

	fmt.Printf("📩 Message: %+v\n", msg)
	return nil
}

// ListAllMessages fetches and prints all messages
func ListAllMessages() error {
	resp, err := client.Get(client.serverURL + "/all")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server responded with: %s", resp.Status)
	}

	var msgs []utils.Message
	if err := json.NewDecoder(resp.Body).Decode(&msgs); err != nil {
		return err
	}

	fmt.Println("🧾 All Messages:")
	for _, msg := range msgs {
		fmt.Printf("- %+v\n", msg)
	}
	return nil
}

// DeleteMessageByID sends a DELETE request for a message by UUID
func DeleteMessageByID(id string) error {
	req, err := http.NewRequest(http.MethodDelete, client.serverURL+"/delete?id="+id, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("delete failed: %s", resp.Status)
	}

	fmt.Println("🗑️ Message deleted.")
	return nil
}
