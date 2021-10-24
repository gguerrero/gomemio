package commands

import (
	"fmt"
)

var ds dataStore

func init() {
	ds = loadDataStore()
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
