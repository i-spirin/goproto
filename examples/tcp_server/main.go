package main

import (
	"log"
	"net"

	"github.com/i-spirin/goproto/tcp_proto"
)

type MyCustomTCPProto struct {
	connection *net.TCPConn
}

func (m *MyCustomTCPProto) ConnectionMade(conn *net.TCPConn) {
	log.Println("Connection created")
	m.connection = conn
}

func (m *MyCustomTCPProto) DataReceived(data []byte) {
	strData := "RESPONSE: " + string(data)
	log.Println("->", string(data))
	_, err := m.connection.Write([]byte(strData))
	if err != nil {
		log.Println("Error sending packet", err)
	}
}

func (m *MyCustomTCPProto) ErrorReceived(e error) {
	log.Printf("Got error: %v\n", e)
}

func (m *MyCustomTCPProto) ConnectionLost(e error) {
	log.Println("Connection closed", e)
}

func main() {

	server := tcp_proto.New(&MyCustomTCPProto{})
	server.Start("tcp4", "0.0.0.0:5088")
	server.Serve()

}
