package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/clavera2/yellow_jacket/utils"
	"github.com/google/uuid"
)

var (
	server = http.NewServeMux()
	mP     = NewMessagePool()
)

func main() {
	initializeServer()
	log.Println("Starting message caching server on :8080")
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func initializeServer() {
	server.HandleFunc("/", homeHandler)
	server.HandleFunc("/add", addMessageHandler)
	server.HandleFunc("/get", getMessageHandler)
	server.HandleFunc("/all", getAllMessagesHandler)
	server.HandleFunc("/delete", deleteMessageHandler)
}

func handleAddMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var msg utils.Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid message payload", http.StatusBadRequest)
		return
	}

	if err := mP.AddMessage(msg); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Message added"))
}

func handleGetMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	msg, err := mP.GetMessage(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(msg)
}

func handleGetAllMessages(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	msgs := mP.GetAllMessages()
	json.NewEncoder(w).Encode(msgs)
}

func handleDeleteMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	if err := mP.DeleteMessage(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Write([]byte("Message deleted"))
}

func handleClearPool(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	mP.ClearPool()
	w.Write([]byte("Message pool cleared"))
}
