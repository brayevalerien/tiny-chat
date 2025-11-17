package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	log.Println("Connecting to chat server...")

	port := ":8000"
	connection, err := net.Dial("tcp", "localhost"+port)
	if err != nil {
		log.Fatal("Failed to connect to server:", err)
	}
	defer connection.Close()

	log.Println("Connected to server!")

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
		fmt.Print("Received: ", message)
	}
}
