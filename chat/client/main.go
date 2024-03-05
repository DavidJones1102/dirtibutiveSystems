package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	connHost = "localhost"
	connPort = "8080"
	connTCP  = "tcp"
	connUDP  = "udp"
)

func main() {
	fmt.Println("Connecting to", connTCP, "server", connHost+":"+connPort)
	socketTCP, err := net.Dial(connTCP, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	fmt.Println("Connecting to", connUDP, "server", connHost+":"+connPort)
	socketUDP, err := net.Dial(connUDP, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	closeChan := make(chan bool)
	go handleOut(socketTCP, closeChan)
	go handleOut(socketUDP, closeChan)
	go handleIn(socketTCP, socketUDP)
	_ = <-closeChan
}
func handleIn(socketTCP net.Conn, socketUDP net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		if strings.HasPrefix(input, "U ") {
			fmt.Println("Sending through UDP")
			socketUDP.Write([]byte(input))
		} else {
			fmt.Println("Sending through TCP")
			socketTCP.Write([]byte(input))
		}
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
