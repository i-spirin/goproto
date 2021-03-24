package tcp_proto

import (
	"io"
	"log"
	"net"
)

type TCPProto struct {
	connection     *net.TCPConn
	listener       *net.TCPListener
	CustomTCPProto CustomTCPProto
}

type CustomTCPProto interface {
	ConnectionMade(*net.TCPConn)
	DataReceived([]byte)
	ErrorReceived(error)
	ConnectionLost(error)
}

func New(c CustomTCPProto) *TCPProto {
	return &TCPProto{CustomTCPProto: c}
}

func (t *TCPProto) Write(b []byte) (n int, err error) {
	return t.connection.Write(b)
}

func (t *TCPProto) Start(network string, address string) error {
	laddr, err := net.ResolveTCPAddr(network, address)
	if err != nil {
		return err
	}
	listener, err := net.ListenTCP(network, laddr)
	if err != nil {
		return err
	}
	t.listener = listener
	return nil
}

func (t *TCPProto) Serve() {
	for {
		conn, err := t.listener.AcceptTCP()
		if err != nil {
			log.Println("Serve error:", err)
			go t.CustomTCPProto.ErrorReceived(err)
			continue
		}
		p := TCPProto{connection: conn, CustomTCPProto: t.CustomTCPProto}
		go p.CustomTCPProto.ConnectionMade(conn)
		go t.HandleConnection(conn)
	}
}

func (t *TCPProto) HandleConnection(connection *net.TCPConn) {
	buffer := make([]byte, 1024)
	for {
		_, err := connection.Read(buffer)
		if err != nil {
			if err != io.EOF {
				go t.CustomTCPProto.ErrorReceived(err)
				continue
			}
			go t.CustomTCPProto.ConnectionLost(err)
			return
		}
		go t.CustomTCPProto.DataReceived(buffer)
	}
}
