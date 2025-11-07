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

	scanner := bufio.NewScanner(os.Stdin) // scans user input
	reader := bufio.NewReader(connection) // read data from server

	for {
		// read user input
		fmt.Print("You: ")
		scanner.Scan()
		message := scanner.Text()

		// send message to the server
		fmt.Fprint(connection, message+"\n")
		response, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Server disconnected:", err)
			break
		}
		fmt.Print("Server: ", response)

	}
}
