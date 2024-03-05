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

func main() {
	fmt.Println("Connecting to", connType, "server", connHost+":"+connPort)
	socket, err := net.Dial(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	go handleOut(socket)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Text to send: ")
		input, _ := reader.ReadString('\n')
		socket.Write([]byte(input))
	}

}

func handleOut(socket net.Conn) {
	for {
		message, err := bufio.NewReader(socket).ReadString('\n')
		if err != nil {
			fmt.Println("I will no longer receive messages")
			return
		}
		log.Println(message)
	}
}
