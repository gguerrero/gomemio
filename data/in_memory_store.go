package data

import "sync"

type InMemoryStore struct {
	data map[string]string
	mu   sync.Mutex
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		data: make(map[string]string),
	}
}

func (ds *InMemoryStore) List() map[string]string {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	return ds.data
}

func (ds *InMemoryStore) Find(key string) string {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	return ds.data[key]
}

func (ds *InMemoryStore) Add(key, value string) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	ds.data[key] = value
}

func (ds *InMemoryStore) Delete(key string) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	delete(ds.data, key)
}
