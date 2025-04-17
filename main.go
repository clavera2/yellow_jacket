package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/clavera2/yellow_jacket/client"
)

func main() {
	//OPENING BANNER FOR YELLOW_JACKET CLIENT
	fmt.Println(":-:-:-:-:-:-:-:-:-:-- WeLcOmE tO yElLoW jAcKeT --:-:-:-:-:-:-:-:-:-:-:")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter server URL (e.g., http://localhost:8080): ")
	serverURL, _ := reader.ReadString('\n')
	serverURL = strings.TrimSpace(serverURL)

	client.InitializeServerURL(serverURL)

	for {
		fmt.Println("\n--- MENU ---")
		fmt.Println("1. Send Message")
		fmt.Println("2. Get Message by ID")
		fmt.Println("3. List All Messages")
		fmt.Println("4. Delete Message by ID")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter message content: ")
			content, _ := reader.ReadString('\n')
			content = strings.TrimSpace(content)

			msg, err := client.MakeMessage(content)
			if err != nil {
				fmt.Println("Error creating message:", err)
				continue
			}

			if err := client.SendMessage(msg); err != nil {
				fmt.Println("Error sending message:", err)
			}

		case "2":
			fmt.Print("Enter message UUID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			if err := client.GetMessageByID(id); err != nil {
				fmt.Println("Error:", err)
			}

		case "3":
			if err := client.ListAllMessages(); err != nil {
				fmt.Println("Error:", err)
			}

		case "4":
			fmt.Print("Enter message UUID to delete: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			if err := client.DeleteMessageByID(id); err != nil {
				fmt.Println("Error:", err)
			}

		case "5":
			fmt.Println("👋 Goodbye from Yellow Jacket!")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
