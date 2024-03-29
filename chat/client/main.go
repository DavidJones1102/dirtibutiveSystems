package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	connHost = "localhost"
	connPort = 8080
	connTCP  = "tcp"
	connUDP  = "udp"
)

func main() {
	fmt.Printf("Connecting to %s server %s:%d\n", connTCP, connHost, connPort)
	socketTCP, err := net.Dial(connTCP, fmt.Sprintf("%s:%d", connHost, connPort))
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	UDPHost, UDPPortString, _ := strings.Cut(socketTCP.LocalAddr().String(), ":")
	UDPPort, _ := strconv.Atoi(UDPPortString)
	socketUDP, err := net.DialUDP(
		connUDP,
		&net.UDPAddr{Port: UDPPort, IP: net.ParseIP(UDPHost)},
		&net.UDPAddr{Port: connPort, IP: net.ParseIP(connHost)},
	)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	closeChan := make(chan bool)
	go handleOut(socketTCP, closeChan)
	go handleOutUDP(*socketUDP, closeChan)
	go handleIn(socketTCP, socketUDP)
	_ = <-closeChan
}
func handleIn(socketTCP net.Conn, socketUDP net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		if strings.HasPrefix(input, "U ") {
			builder := strings.Builder{}
			builder.WriteString(input)
			for {
				in, _ := reader.ReadString('\n')
				if len(in) == 2 {
					break
				}
				builder.WriteString(in)
			}
			fmt.Println("Sending through UDP")
			socketUDP.Write([]byte(builder.String()))
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

func handleOutUDP(socket net.UDPConn, close chan<- bool) {
	for {
		message := make([]byte, 2048)
		socket.ReadFrom(message)
		log.Printf("Got msg through UDP: %s", message)
	}
}
