package main

import (
	"fmt"
	"log"
	"net/http"

	server "github.com/clavera2/yellow_jacket/server/internals"
)

func main() {
	fmt.Println("🚀 Starting Yellow-Jacket Server...")

	// Initialize all routes and handler functions
	server.Initialize()
	log.Fatal(http.ListenAndServe(":8080", server.Router()))

	// Start the server
	addr := ":8080"
	fmt.Printf("🟢 Server running at http://localhost%s\n", addr)

	if err := http.ListenAndServe(addr, server.Router()); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}
