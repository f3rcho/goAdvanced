package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

type Client chan<- string

var (
	incomingClients = make(chan Client)
	leavingClients  = make(chan Client)
	clientMessages  = make(chan string)
)

var (
	host = flag.String("h", "localhost", "host")
	port = flag.Int("p", 3090, "port")
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	message := make(chan string)

	go MessageWriter(conn, message)

	clientName := conn.RemoteAddr().String()

	message <- fmt.Sprintf("Welcome to the server, your name %s\n", clientName)
	clientMessages <- fmt.Sprintf("New client is here, name %s\n", clientName)
	incomingClients <- message

	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		clientMessages <- fmt.Sprintf("%s: %s\n", clientName, inputMessage.Text())
	}

	leavingClients <- message
	clientMessages <- fmt.Sprintf("%s said goodbye!", clientName)
}

func MessageWriter(conn net.Conn, messages <-chan string) {
	for message := range messages {
		fmt.Fprintln(conn, message)
	}
}

func Broadcast() {
	clients := make(map[Client]bool)

	for {
		select {
		case message := <-clientMessages:
			for client := range clients {
				client <- message
			}
		case newClient := <-incomingClients:
			clients[newClient] = true
		case leavingClient := <-leavingClients:
			delete(clients, leavingClient)
			close(leavingClient)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatalf("Listener error:%v", err)
	}
	go Broadcast()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("listener connection error:%v", err)
			continue
		}
		go HandleConnection(conn)
	}
}
