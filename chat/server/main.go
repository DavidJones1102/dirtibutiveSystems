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
	connPort = 8080
	connTCP  = "tcp"
	connUDP  = "udp"
)

func main() {
	fmt.Printf("Starting %s server on %s:%d", connTCP, connHost, connPort)
	socketTCP, err := net.Listen(connTCP, fmt.Sprintf("%s:%d", connHost, connPort))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer socketTCP.Close()
	socketUDP, err := net.ListenUDP(connUDP, &net.UDPAddr{Port: connPort, IP: net.ParseIP(connHost)})
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer socketUDP.Close()
	go readUDP(socketUDP)

	messageChannel := make(chan pkg.Message, 2)
	addClientChannel := make(chan pkg.Client, 2)
	removeClientChannel := make(chan net.Addr, 2)
	go distributeMessages(messageChannel, addClientChannel, removeClientChannel)

	for {
		client, err := socketTCP.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}

		fmt.Println("Client connected.")
		fmt.Println("Client " + client.RemoteAddr().String() + " connected.")
		handleClient(client, messageChannel, addClientChannel, removeClientChannel)
	}
}

func handleClient(conn net.Conn, messageChannel chan<- pkg.Message,
	addClientChannel chan<- pkg.Client,
	removeClientChannel chan<- net.Addr) {
	channel := make(chan string)
	client := pkg.CreateClient(conn, channel)
	addClientChannel <- client
	go handleClientIn(conn, messageChannel, removeClientChannel)
	go handleClientOut(client)
}

func handleClientIn(conn net.Conn, messageChannel chan<- pkg.Message, removeClientChannel chan<- net.Addr) {
	for {
		buffer, err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil {
			conn.Close()
			removeClientChannel <- conn.RemoteAddr()
			return
		}
		messageChannel <- pkg.CreateMessage(conn.RemoteAddr(), string(buffer))
	}
}
func handleClientOut(client pkg.Client) {
	for msg := range client.GetOutChan() {
		client.Send(msg)
	}
}

func distributeMessages(messageChannel <-chan pkg.Message, addClientChan <-chan pkg.Client, removeClientChannel <-chan net.Addr) {
	var clientsMap = make(map[net.Addr]chan<- string, 2)
	for {
		select {
		case clientMsg := <-messageChannel:
			for addr, c := range clientsMap {
				if addr == clientMsg.GetAddr() {
					continue
				}
				c <- "msg from " + clientMsg.GetStrAddr() + " - " + clientMsg.GetMsg()
			}
		case client := <-addClientChan:
			clientsMap[client.GetAddr()] = client.GetOutChan()
			fmt.Printf("Client %s added!\n", client.GetAddr().String())
		case addr := <-removeClientChannel:
			delete(clientsMap, addr)
			fmt.Printf("Client %s removed!\n", addr)
		}
	}
}

func readUDP(connUDP *net.UDPConn) {
	for {
		message := make([]byte, 2048)
		_, remoteAddr, err := connUDP.ReadFromUDP(message)
		if err != nil {
			fmt.Println("Closing...")
			return
		}
		log.Printf("%s - %s", remoteAddr, message)
	}
}
