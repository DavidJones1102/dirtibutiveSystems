package pkg

import (
	"net"
)

// Client structure
type Client struct {
	conn    net.Conn
	outChan chan string
}

func CreateClient(conn net.Conn, outChan chan string) Client {
	return Client{conn, outChan}
}

func (c Client) GetAddr() net.Addr {
	return c.conn.RemoteAddr()
}
func (c Client) GetOutChan() chan string {
	return c.outChan
}
func (c Client) Send(msg string) {
	c.conn.Write([]byte(msg))
}

// Message structure
type Message struct {
	addr net.Addr
	msg  string
}

func CreateMessage(addr net.Addr, msg string) Message {
	return Message{addr, msg}
}

func (m Message) GetStrAddr() string {
	return m.addr.String()
}
func (m Message) GetAddr() net.Addr {
	return m.addr
}
func (m Message) GetMsg() string {
	return m.msg
}
