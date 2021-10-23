package service

import (
	"log"
	"net"
	"os"
	"os/signal"
)

type handler struct {
	listener    net.Listener
	connections []*connection
}

func NewHandler(li net.Listener) *handler {
	return &handler{listener: li}
}

func (h *handler) Handle() error {
	// defer h.closeAndLog()

	h.handleOSInterrupt()

	for {
		conn, err := h.listener.Accept()
		if err != nil {
			return err
		}

		log.Println("Connection accepted from", conn.RemoteAddr())
		c := NewConnection(conn)

		h.connections = append(h.connections, c)
		// connIndex := len(h.connections) - 1
		go h.handleConn(c)
	}
}

func (h *handler) closeAndLog() {
	log.Printf("Stopping listener (%s)\n", h.listener.Addr())
	for _, c := range h.connections {
		c.closeAndLog()
	}
	h.listener.Close()
}

func (h *handler) handleConn(c *connection) {
	if err := c.handle(); err != nil {
		log.Println(err)
	}
}

// func (h *handler) deleteConn(i int) []*connection {
// 	last := len(h.connections) - 1

// 	h.connections[i] = h.connections[last]
// 	return h.connections[:last]
// }

func (h *handler) handleOSInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for sig := range c {
			if sig.String() == "interrupt" {
				log.Println("... ^C captured, stopping server!")
				h.closeAndLog()
				os.Exit(0)
			}
		}
	}()
}
