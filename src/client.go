package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func promptForUsername() string {
	fmt.Print("Choose a username: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	username := scanner.Text()
	if username == "" {
		log.Fatal("Username cannot be empty.")
	}
	return username
}

func main() {
	username := promptForUsername()

	log.Println("Connecting to chat server...")

	port := ":8000"
	connection, err := net.Dial("tcp", "localhost"+port)
	if err != nil {
		log.Fatal("Failed to connect to server:", err)
	}
	defer connection.Close()

	log.Println("Connected to server!")

	// server expects username as first message
	_, err = fmt.Fprintf(connection, "%s\n", username) // server uses newline as delimiter
	if err != nil {
		log.Fatal("Failed to send username to server:", err)
	}

	go receiveMessages(connection)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("You: ")
		if !scanner.Scan() {
			break
		}
		message := scanner.Text()
		_, err := fmt.Fprint(connection, message+"\n")
		if err != nil {
			log.Println("Failed to send message:", err)
			os.Exit(1)
		}
	}
}

func receiveMessages(connection net.Conn) {
	reader := bufio.NewReader(connection)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Server disconnected:", err)
			os.Exit(0)
		}
		fmt.Print("\r\033[K" + message + "You: ")
	}
}
