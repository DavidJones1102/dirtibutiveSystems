package main

import (
	"bufio"
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

var clientsMap = make(map[net.Addr]chan<- string)

func main() {
	fmt.Println("Starting " + connType + " server on " + connHost + ":" + connPort)
	socket, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer socket.Close()

	for {
		client, err := socket.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}
		channel := make(chan string)
		clientsMap[client.RemoteAddr()] = channel
		fmt.Println("Client connected.")
		fmt.Println("Client " + client.RemoteAddr().String() + " connected.")
		fmt.Println("Client " + client.LocalAddr().String() + " connected.")

		go handleConnection(client, channel)
		propagateMessage(client.RemoteAddr(), "-----")
	}
}

func handleConnection(conn net.Conn, c <-chan string) {
	log.Println("Waiting for message")
	select {
	case msg := <-c:
		conn.Write([]byte(msg))

	}
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		fmt.Println("Client left.")
		conn.Close()
		return
	}

	log.Println("Client message:", string(buffer[:len(buffer)-1]))

	conn.Write(buffer)
	handleConnection(conn, c)
}

func propagateMessage(senderAddr net.Addr, msg string) {
	for addr, c := range clientsMap {
		if addr == senderAddr {
			continue
		}
		c <- "msg from " + senderAddr.String() + " - " + msg
	}
}
