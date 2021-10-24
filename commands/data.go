package commands

import "github.com/gguerrero/gomemio/data"

// dataStore interface abstracts the boundary with the data layer. The data layer should implement this interface.
type dataStore interface {
	List() map[string]string
	Find(key string) string
	Add(key, value string)
	Delete(key string)
}

func loadDataStore() dataStore {
	// Returns the only dataStore implementation, we might want to choose the impl. in the future
	return data.NewInMemoryStore()
}
