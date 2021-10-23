package data

import "sync"

type DataStore struct {
	data map[string]string
	mu   sync.Mutex
}

func NewDataStore() *DataStore {
	return &DataStore{
		data: make(map[string]string),
	}
}

func (ds *DataStore) List() map[string]string {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	return ds.data
}

func (ds *DataStore) Find(key string) string {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	return ds.data[key]
}

func (ds *DataStore) Add(key, value string) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	ds.data[key] = value
}

func (ds *DataStore) Delete(key string) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	delete(ds.data, key)
}
