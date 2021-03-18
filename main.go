package main

import (
	"log"
	"net"

	"github.com/i-spirin/goproto/udp_proto"
)

func main() {

	var connection *net.UDPConn

	ConnectionMade := func(conn *net.UDPConn) {
		log.Println("Connection created")
		connection = conn
	}

	DatagramReceived := func(addr *net.UDPAddr, data []byte) {
		strData := string(data) + "\n"
		log.Println(addr, "->", string(data))
		_, err := connection.WriteToUDP([]byte(strData), addr)
		if err != nil {
			log.Println("Error sending packet", err)
		}
	}

	ErrorReceived := func(e error) {
		log.Printf("Got error: %v\n", e)
	}

	ConnectionLost := func(e error) {
		log.Println("Connection closed", e)
	}

	server := udp_proto.New(ConnectionMade, DatagramReceived, ErrorReceived, ConnectionLost)
	server.Start("udp4", "0.0.0.0", 5088)
	server.Serve()

}
