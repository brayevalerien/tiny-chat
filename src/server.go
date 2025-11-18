package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

// maps net.Conn to the usernames
var clients map[net.Conn]string
var clientsMutex sync.Mutex

func main() {
	log.Println("Starting chat server...")

	clients = make(map[net.Conn]string)

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

	reader := bufio.NewReader(connection)

	username, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Failed to read username:", err)
		return
	}
	username = strings.TrimSpace(username)
	clientsMutex.Lock()
	clients[connection] = username
	clientsMutex.Unlock()
	defer removeClient(connection)

	broadcastAll(username + " joined the chat\n")
	log.Println("New client connected:", username, "- Current connections:", len(clients))

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Client disconnected:", err)
			return
		}

		clientsMutex.Lock()
		currentUsername := clients[connection]
		clientsMutex.Unlock()
		message = strings.TrimSpace(message)
		formattedMessage := fmt.Sprintf("[%s] %s\n", currentUsername, message)

		log.Println(formattedMessage)

		broadcastAll(formattedMessage)
	}
}

func broadcastAll(message string) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()
	for client := range clients {
		fmt.Fprint(client, message)
	}
}

func removeClient(connection net.Conn) {
	clientsMutex.Lock()
	username := clients[connection]
	delete(clients, connection)
	clientsMutex.Unlock()

	broadcastAll(username + " left the chat\n")
}
