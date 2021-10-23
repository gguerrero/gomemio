package gomemio

import (
	"fmt"
	"log"
	"net"

	"github.com/gguerrero/gomemio/service"
)

type port uint16
type server struct {
	address net.IP
	port
}

const (
	network = "tcp"
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

	return service.NewHandler(li).Handle()
}

func (s *server) String() string {
	return fmt.Sprintf("%s:%d", s.address, s.port)
}
