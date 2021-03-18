package udp_proto

import (
	"fmt"
	"net"
)

type UDPProto struct {
	connection       *net.UDPConn
	ConnectionMade   func(*net.UDPConn)
	DatagramReceived func(*net.UDPAddr, []byte)
	ErrorReceived    func(error)
	ConnectionLost   func(error)
}

func New(
	ConnectionMade func(*net.UDPConn),
	DatagramReceived func(*net.UDPAddr, []byte),
	ErrorReceived func(error),
	ConnectionLost func(error)) *UDPProto {
	return &UDPProto{
		ConnectionMade:   ConnectionMade,
		DatagramReceived: DatagramReceived,
		ErrorReceived:    ErrorReceived,
		ConnectionLost:   ConnectionLost,
	}
}

func (u *UDPProto) Write(p []byte, addr *net.UDPAddr) (n int, err error) {
	return u.connection.WriteToUDP(p, addr)
}

func (u *UDPProto) Start(network string, host string, port int) error {
	s, err := net.ResolveUDPAddr(network, fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return fmt.Errorf("error resolving address: %v", err)
	}

	u.connection, err = net.ListenUDP("udp", s)
	if err != nil {
		return fmt.Errorf("error binding a port: %v", err)
	}

	go u.ConnectionMade(u.connection)
	return nil
}

func (u *UDPProto) Serve() {
	buffer := make([]byte, 1024)

	for {
		n, addr, _ := u.connection.ReadFromUDP(buffer)
		go u.DatagramReceived(addr, buffer[0:n-1])
	}

}

func (u *UDPProto) Close(e error) {
	u.ConnectionLost(e)
}
