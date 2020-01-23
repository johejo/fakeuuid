package fakeuuid

import (
	"sync"

	"github.com/google/uuid"
)

var (
	mutex   sync.RWMutex
	backend string
)

type UUID struct{}

// String returns string from backend. If backend is empty returns a new real uuid.New().String().
func (u UUID) String() string {
	if b := getBackend(); b != "" {
		return b
	}
	return uuid.New().String()
}

func getBackend() string {
	mutex.Lock()
	defer mutex.Unlock()
	return backend
}

// Mock mocks fakeuuid.New().String(), and returns restore callback.
func Mock(fake string) func() {
	mutex.Lock()
	defer mutex.Unlock()
	origin := backend
	backend = fake
	return func() { Mock(origin) }
}

// New returns a New fakeuuid.UUID.
func New() UUID {
	return UUID{}
}
