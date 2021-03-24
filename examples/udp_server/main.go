// package main

// import (
// 	"log"
// 	"net"

// 	"github.com/i-spirin/goproto/udp_proto"
// )

// type MyCustomUDPProto struct {
// 	connection *net.UDPConn
// }

// func (m *MyCustomUDPProto) ConnectionMade(conn *net.UDPConn) {
// 	log.Println("Connection created")
// 	m.connection = conn
// }

// func (m *MyCustomUDPProto) DatagramReceived(addr *net.UDPAddr, data []byte) {
// 	strData := string(data) + "\n"
// 	log.Println(addr, "->", string(data))
// 	_, err := m.connection.WriteToUDP([]byte(strData), addr)
// 	if err != nil {
// 		log.Println("Error sending packet", err)
// 	}
// }

// func (m *MyCustomUDPProto) ErrorReceived(e error) {
// 	log.Printf("Got error: %v\n", e)
// }

// func (m *MyCustomUDPProto) ConnectionLost(e error) {
// 	log.Println("Connection closed", e)
// }

// func main() {

// 	server := udp_proto.New(&MyCustomUDPProto{})
// 	server.Start("udp4", "0.0.0.0", 5088)
// 	server.Serve()

// }
