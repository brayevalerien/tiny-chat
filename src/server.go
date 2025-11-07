package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	log.Println("Starting chat server...")

	port := ":8000"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
	defer listener.Close()

	log.Println("Server listening on", port)

	connection, err := listener.Accept()
	if err != nil {
		log.Fatal("Failed to accept client connection:", err)
	}
	defer connection.Close()

	log.Println("Client connected!")

	reader := bufio.NewReader(connection) // read from client

	for {
		clientMessage, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Client disconnected:", err)
			break
		}
		log.Println("Client:", clientMessage)

		fmt.Fprint(connection, clientMessage)
	}
}
