package commands

import (
	"fmt"
	"strings"
)

type action uint8

const (
	GET = iota
	SET
	DEL
	LIST
	EXIT
	UNKNOWN
)

type command struct {
	action
	key   string
	value string
}

var ds dataStore

func init() {
	ds = loadDataStore()
}

func NewCommand(cmd []string) (*command, error) {
	action, err := parseAction(cmd[0])
	if err != nil {
		return &command{}, err
	}

	var key, value string

	if len(cmd) > 1 {
		key = cmd[1]
	}

	if len(cmd) > 2 {
		value = strings.Join(cmd[2:], " ")
	}

	return &command{
		action: action,
		key:    key,
		value:  value,
	}, nil
}

func parseAction(a string) (action, error) {
	switch strings.ToUpper(a) {
	case "GET":
		return GET, nil
	case "SET":
		return SET, nil
	case "DEL":
		return DEL, nil
	case "LIST":
		return LIST, nil
	case "EXIT":
		return EXIT, nil
	default:
		return UNKNOWN, fmt.Errorf("commands parseAction: Unknown action %s", a)
	}
}

func (cmd *command) IsExit() bool {
	return cmd.action == EXIT
}

func (cmd *command) Execute() (string, error) {
	switch cmd.action {
	case GET:
		return ds.Find(cmd.key), nil
	case SET:
		ds.Add(cmd.key, cmd.value)
		return "OK", nil
	case DEL:
		ds.Delete(cmd.key)
		return "DELETED", nil
	case LIST:
		data := ds.List()
		return fmt.Sprint(data), nil
	default:
		return "", fmt.Errorf("commands excute: cannot excute action %d", cmd.action)
	}
}
