package service

import (
	"fmt"
	"log"
	"net"

	"github.com/gguerrero/gomemio/commands"
	"github.com/gguerrero/gomemio/scanner"
)

type connection struct {
	net.Conn
}

func NewConnection(conn net.Conn) *connection {
	return &connection{conn}
}

func (conn *connection) handle() error {
	defer conn.closeAndLog()

	s := scanner.NewScanner(conn)
	go s.ScanLines()

	conn.parseCommands(s.Commands)

	return nil
}

func (conn *connection) closeAndLog() {
	log.Println("Closing connection from", conn.RemoteAddr())
	conn.Close()
}

func (conn *connection) parseCommands(commandCh <-chan []string) {
	for command := range commandCh {
		cmd, err := commands.NewCommand(command)
		if err != nil {
			log.Println(err)
			fmt.Fprintln(conn, err)
			continue
		}

		if cmd.IsExit() {
			break
		}

		result, err := cmd.Execute()
		if err != nil {
			log.Println(err)
			fmt.Fprintln(conn, err)
			continue
		}

		fmt.Fprintln(conn, result)
	}
}
