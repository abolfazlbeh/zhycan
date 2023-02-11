package http

import (
	"encoding/json"
	"log"
	"sync"
	"zhycan/internal/config"
)

// Mark: manager

// manager object
type manager struct {
	name    string
	lock    sync.Mutex
	servers []*Server
}

// MARK: Module variables
var managerInstance *manager = nil
var once sync.Once

// Module init function
func init() {
	log.Println("HTTP Manager Package Initialized...")
}

// init - Manager Constructor - It initializes the manager configuration params
func (m *manager) init() {
	m.name = "http"

	m.lock.Lock()
	defer m.lock.Unlock()

	// read configs and save it
	serversCfg, err := config.GetManager().Get(m.name, "servers")
	if err != nil {
		return
	}

	for _, item := range serversCfg.([]interface{}) {
		jsonBody, err2 := json.Marshal(item)
		if err2 != nil {
			continue
		}

		obj := ServerConfig{}
		err := json.Unmarshal(jsonBody, &obj)
		if err == nil {
			server, err1 := NewServer(obj)
			if err1 == nil {
				m.servers = append(m.servers, server)
			}
		}
	}
}

// restartOnChangeConfig - subscribe a function for when the config is changed
func (m *manager) restartOnChangeConfig() {
	m.lock.Lock()
	defer m.lock.Unlock()

	// Config config server to reload
	wrapper, err := config.GetManager().GetConfigWrapper(m.name)
	if err == nil {
		wrapper.RegisterChangeCallback(func() interface{} {
			m.StopServers()
			m.init()
			m.StartServers()
			return nil
		})
	} else {
		// TODO: make some logs
	}
}

// MARK: Public Functions

// GetManager - This function returns singleton instance of Logger Manager
func GetManager() *manager {
	// once used for prevent race condition and manage critical section.
	once.Do(func() {
		managerInstance = &manager{}
		managerInstance.init()
		managerInstance.restartOnChangeConfig()
	})
	return managerInstance
}

// StartServers - iterate over all servers and start them
func (m *manager) StartServers() error {
	for _, item := range m.servers {
		go func(s *Server) {
			err := s.Start()
			if err != nil {
				// TODO: print some error
			}
		}(item)
	}
	return nil
}

// StopServers - iterate over all severs and stop them
func (m *manager) StopServers() error {
	return NewNotImplementedErr()
}
