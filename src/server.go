package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
)

var clients []net.Conn      // stores connected clients
var clientsMutex sync.Mutex // mutex on clients

func main() {
	log.Println("Starting chat server...")

	port := ":8000"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
	defer listener.Close()

	log.Println("Server listening on", port)

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection:", err)
			continue
		}
		go handleNewClient(connection)
	}
}

func handleNewClient(connection net.Conn) {
	defer connection.Close()
	clientsMutex.Lock()
	clients = append(clients, connection)
	clientsMutex.Unlock()
	defer removeClient(connection)
	log.Println("New client connected. Current number of connections:", len(clients))
	reader := bufio.NewReader(connection)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Client disconnected:", err)
			return
		}
		log.Println("Client message:", message)
		broadcast(message, connection)
	}
}

// sends a message to all clients other than sender
func broadcast(message string, sender net.Conn) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()
	for _, receiver := range clients {
		if receiver != sender {
			fmt.Fprint(receiver, message)
		}
	}
}

func removeClient(connection net.Conn) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()
	for i, client := range clients {
		if client == connection {
			clients = append(clients[:i], clients[i+1:]...)
			return
		}
	}
}
