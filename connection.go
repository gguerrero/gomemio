package gomemio

import (
	"net"
	"time"
)

type connection struct {
	net.Conn
}

func NewConnection(conn net.Conn) *connection {
	return &connection{conn}
}

func (conn *connection) Handle() error {
	defer conn.Close()

	conn.SetDeadline(time.Now().Add(timeout))

	return nil
}
