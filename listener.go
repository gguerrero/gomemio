package gomemio

import (
	"log"
	"net"
)

type listener struct {
	net.Listener
}

func NewListener(li net.Listener) *listener {
	return &listener{li}
}

func (li *listener) Handle() error {
	defer li.closeAndLog()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		go func() {
			err = NewConnection(conn).Handle()
			if err != nil {
				log.Println(err)
			}
		}()
	}

	return nil
}

func (li *listener) closeAndLog() {
	log.Printf("Stopping listener (%s)\n", li.Addr())
	li.Close()
}
