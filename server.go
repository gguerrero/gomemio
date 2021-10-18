package gomemio

import (
	"fmt"
	"log"
	"net"
	"time"
)

type port uint16
type server struct {
	address net.IP
	port
}

const (
	network = "tcp"
	// TCP requests default timeout.
	timeout = time.Second * 10
)

func NewServer(a string, p int) *server {
	return &server{
		address: net.ParseIP(a),
		port:    port(p),
	}
}

func (s *server) ListenAndServe() error {
	log.Printf("MemIO listening at %s\n", s.String())

	li, err := net.Listen(network, s.String())
	if err != nil {
		return err
	}

	return NewListener(li).Handle()
}

func (s *server) String() string {
	return fmt.Sprintf("%s:%d", s.address, s.port)
}
