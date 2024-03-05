package main

import (
	"bufio"
	"chat/server/pkg"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
)

func main() {
	fmt.Println("Starting " + connType + " server on " + connHost + ":" + connPort)
	socket, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer socket.Close()

	messageChannel := make(chan pkg.Message, 2)
	addClientChannel := make(chan pkg.Client, 2)
	go propagateMessage(messageChannel, addClientChannel)
	messageChannel <- pkg.CreateMessage(nil, "Test")

	for {
		client, err := socket.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}

		fmt.Println("Client connected.")
		fmt.Println("Client " + client.RemoteAddr().String() + " connected.")

		handleClient(client, messageChannel, addClientChannel)
	}
}

func handleClient(client net.Conn, messageChannel chan<- pkg.Message, addClientChannel chan<- pkg.Client) {
	channel := make(chan string)
	addClientChannel <- pkg.CreateClient(client, channel)
	go handleClientIn(client, messageChannel)
	go handleClientOut(client, channel)
}

func handleClientIn(conn net.Conn, messageChannel chan<- pkg.Message) {
	fmt.Println("I will send in a moment")
	for {
		buffer, err := bufio.NewReader(conn).ReadBytes('\n')

		if err != nil {
			fmt.Println("Client left.")
			conn.Close()
			return
		}

		log.Println("Client message:", string(buffer[:len(buffer)-1]))
		messageChannel <- pkg.CreateMessage(conn.RemoteAddr(), string(buffer))
		log.Println("Message passed:", string(buffer[:len(buffer)-1]))
	}

}
func handleClientOut(conn net.Conn, c <-chan string) {
	fmt.Println("Waiting for messages to me!")
	for msg := range c {
		fmt.Println("Sending to client")
		conn.Write([]byte(msg))
		fmt.Println("Sent to client")
	}
}

func propagateMessage(messageChannel <-chan pkg.Message, addClientChan <-chan pkg.Client) {
	var clientsMap = make(map[net.Addr]chan<- string, 2)
	select {
	case clientMsg := <-messageChannel:
		fmt.Println("Got a message to propagate")
		for addr, c := range clientsMap {
			if addr == clientMsg.GetAddr() {
				continue
			}
			c <- "msg from " + clientMsg.GetStrAddr() + " - " + clientMsg.GetMsg()
		}
	case client := <-addClientChan:
		clientsMap[client.GetAddr()] = client.GetOutChan()
		fmt.Printf("Client %s Added!\n", client.GetAddr().String())
		client.GetOutChan() <- "Server response"
	}
}
