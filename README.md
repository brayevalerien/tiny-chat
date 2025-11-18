# tiny-chat
A tiny multi-client TCP chat server and client implemented in Go for learning purposes.

This is a toy implementation designed to understand TCP networking fundamentals and Go's `net` package. It demonstrates connection handling, goroutines for concurrency, and thread-safe state management with mutexes.

## Running the project

Start the server first.
```bash
go run src/server.go
```

Then start as many new clients you wish by opening new terminals and running the following.
```bash
go run src/client.go
```

Each client will prompt for a username before connecting, usernames are used during the entire session. Restart the client to change username.

## Architecture
Server handles:
- accepting multiple client connections
- reading username as first message
- broadcasting messages to all connected clients

Client handles:
- username input before connecting
- sending and receiving messages