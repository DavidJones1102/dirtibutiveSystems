package main

import (
	"bufio"
	"chat/server/pkg"
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

	messageTCPChannel := make(chan pkg.Message, 2)
	addClientChannel := make(chan pkg.Client, 2)
	removeClientChannel := make(chan net.Addr, 2)
	go readUDP(socketUDP, messageTCPChannel)
	go distributeMessages(socketUDP, messageTCPChannel, addClientChannel, removeClientChannel)

	for {
		client, err := socketTCP.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}

		fmt.Println("Client connected.")
		fmt.Println("Client " + client.RemoteAddr().String() + " connected.")
		handleClient(client, messageTCPChannel, addClientChannel, removeClientChannel)
	}
}

func handleClient(conn net.Conn, messageTCPChannel chan<- pkg.Message,
	addClientChannel chan<- pkg.Client,
	removeClientChannel chan<- net.Addr) {
	channel := make(chan string)
	client := pkg.CreateClient(conn, channel)
	addClientChannel <- client
	go handleClientIn(conn, messageTCPChannel, removeClientChannel)
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

func distributeMessages(socketUDP *net.UDPConn, messageChannel <-chan pkg.Message, addClientChan <-chan pkg.Client, removeClientChannel <-chan net.Addr) {
	var clientsMap = make(map[net.Addr]chan<- string, 2)
	for {
		select {
		case clientMsg := <-messageChannel:
			for addr, c := range clientsMap {
				if addr.String() == clientMsg.GetAddr().String() {
					continue
				}
				msg := fmt.Sprintf("msg from %s - %s\n", clientMsg.GetStrAddr(), clientMsg.GetMsg())
				if strings.HasPrefix(clientMsg.GetMsg(), "U ") {
					_, err := socketUDP.WriteToUDP([]byte(msg), stringToUDPAddr(addr.String()))
					if err != nil {
						return
					}
					println("Sending UDP " + addr.String())
				} else {
					c <- msg
				}
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

func readUDP(connUDP *net.UDPConn, messageChannel chan<- pkg.Message) {
	for {
		message := make([]byte, 2048)
		_, remoteAddr, err := connUDP.ReadFromUDP(message)
		if err != nil {
			fmt.Println("Closing...")
			return
		}
		messageChannel <- pkg.CreateMessage(remoteAddr, string(message))
		log.Printf("%s - %s", remoteAddr, message)
	}
}

func stringToUDPAddr(addr string) *net.UDPAddr {
	UDPHost, UDPPortString, _ := strings.Cut(addr, ":")
	UDPPort, _ := strconv.Atoi(UDPPortString)
	return &net.UDPAddr{Port: UDPPort, IP: net.ParseIP(UDPHost)}
}
