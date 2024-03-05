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
	closeChan := make(chan bool)
	go handleOut(socket, closeChan)
	go handleIn(socket)
	_ = <-closeChan
}
func handleIn(socket net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		socket.Write([]byte(input))
	}
}
func handleOut(socket net.Conn, close chan<- bool) {
	for {
		message, err := bufio.NewReader(socket).ReadString('\n')
		if err != nil {
			fmt.Println("Closing...")
			close <- true
			return
		}
		log.Println(message)
	}
}
