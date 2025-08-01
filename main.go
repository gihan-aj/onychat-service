package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// Client represents a single connected user. We use a channel for outgoing messages
type client chan<- string

var (
	// entering is a channel for new clients connecting
	entering = make(chan client)
	// leaving is a channel for clients disconnecting.
	leaving = make(chan client)
	// messages is the main channel to braodcast messages to all clients.
	messages = make(chan string)
)

// braodcaster is the core of our chat server.
// It listens on the three channels and manages the set of active clients.
func braodcaster() {
	// clients is a map that holds all currently connected clients.
	// The map's value is 'true', but we don't actually use it. We're just using the map as a set.
	clients := make(map[client]bool)

	for {
		// The 'select' statment blocks until one of its cases can run.
		// This let's us handle multiple channel operations concurrently.
		select {
		case msg := <-messages:
			// When a message is received on the 'messages' channel...
			// ... broadcast it to all connected clients.
			log.Printf("Braodcasting message to %d clients: %s", len(clients), msg)
			for cli := range clients {
				cli <- msg // Send the message to the client's channel
			}

		case cli := <-entering:
			// When a new client connects, add them to our set of clients.
			clients[cli] = true

		case cli := <-leaving:
			// When the client disconnects, remove them from the set...
			delete(clients, cli)
			// ...and close their channel to signal that no more messages will be sent.
			close(cli)
		}
	}
}

// clientWriter is responsible for writing messages from the broadcaster to the client's connection
func clientWriter(conn net.Conn, ch <-chan string){
	for msg := range ch {
		fmt.Fprintln(conn, msg) // Fprintln adds a newline, which is helpful for clients like telnet.
	}
}

// handleConnection now manages a new client's lifecycle.
func handleConnection(conn net.Conn) {
	// Create a channel for this client's outgoing messages.
	ch := make(chan string)
	// Start a new goroutine that will write messages from the channel to the connection
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who // Send a welcome message to the new client.
	messages <- who + " has arrived" // Announce the new client to everyone else.
	entering <- ch // Add this new client to the broadcaster's list

	// Read incoming messages from the client
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	// When the loop finishes (e.g. client dissconnects), announce their departure.
	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func main() {
	// Define the address and port for the server to listen on.
	port := ":8080"

	// Start listening for incoming TCP conections on the specified port.
	listener, err := net.Listen("tcp", port)
	if err != nil {
		// If there's an error (e.g., the port is already in use), log it and exit.
		log.Fatalf("Failed to start server: %v", err);
	}

	// Ensure the listener is closed when the main function exits.
	defer listener.Close()

	// Start  the broadcaster in  its own goroutine so it can run in the background.
	go braodcaster()

	log.Printf("Server listening on port %s", port)

	// Loop forever, waiting for and accepting new client connections.
	for {
		// Accept() blocks until a new connection is made, then returns it.
		conn, err := listener.Accept()
		if err != nil {
			// If there is an error accepting a new connection, log it and continue.
			log.Printf("Failed to accept connection %v", err)
			continue
		}

		// Handle each connection concurrently in its own 'goroutine'.
		// This is the magic of Go's concurrency, The 'go' keyword starts a new goroutine.
		// It allows the server to handle many clients at the same time without waiting
		go handleConnection(conn)
	}
}