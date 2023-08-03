package db

import (
	"log"
	"sync"
)

// Mark: manager

// manager object
type manager struct {
	name string
	lock sync.Mutex
}

// MARK: Module variables
var managerInstance *manager = nil
var once sync.Once

// Module init function
func init() {
	log.Println("DB Manager Package Initialized...")
}

func (m *manager) init() {
	m.name = "db"

	m.lock.Lock()
	defer m.lock.Unlock()
}
