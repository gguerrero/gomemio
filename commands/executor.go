package commands

import (
	"fmt"

	"github.com/gguerrero/gomemio/data"
)

var dataStore *data.DataStore

func init() {
	dataStore = data.NewDataStore()
}

func (cmd *command) Execute() (string, error) {
	switch cmd.action {
	case GET:
		return dataStore.Find(cmd.key), nil
	case SET:
		dataStore.Add(cmd.key, cmd.value)
		return "OK", nil
	case DEL:
		dataStore.Delete(cmd.key)
		return "DELETED", nil
	case LIST:
		data := dataStore.List()
		return fmt.Sprint(data), nil
	default:
		return "", fmt.Errorf("commands excute: cannot excute action %d", cmd.action)
	}
}
