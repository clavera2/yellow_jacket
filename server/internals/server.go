package server

import (
	"encoding/json"
	"net/http"

	"github.com/clavera2/yellow_jacket/utils"
	"github.com/google/uuid"
)

var (
	router = http.NewServeMux()
	mP     = NewMessagePool()
)

func Initialize() {
	initializeServer()
}

func Router() *http.ServeMux {
	return router
}

// initializeServer initializes all routes on the router
func initializeServer() {
	//Register handlers
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/add", handleAddMessage)
	router.HandleFunc("/get", handleGetMessage)
	router.HandleFunc("/all", handleGetAllMessages)
	router.HandleFunc("/delete", handleDeleteMessage)
}

// handleAddMessage handles adding a new message
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

// handleGetMessage retrieves a message by ID
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

// handleGetAllMessages handles listing all messages
func handleGetAllMessages(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	msgs := mP.GetAllMessages()
	json.NewEncoder(w).Encode(msgs)
}

// handleDeleteMessage handles deleting a message by ID
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
