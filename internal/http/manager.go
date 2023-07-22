package http

import (
	"encoding/json"
	"github.com/abolfazlbeh/zhycan/internal/config"
	"github.com/abolfazlbeh/zhycan/internal/utils"
	"github.com/gofiber/fiber/v2"
	"log"
	"sync"
)

// Mark: manager

// manager object
type manager struct {
	name          string
	lock          sync.Mutex
	servers       map[string]*Server
	defaultServer string
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

	var serverNames []string
	m.servers = make(map[string]*Server)

	for _, item := range serversCfg.([]interface{}) {
		jsonBody, err2 := json.Marshal(item)
		if err2 != nil {
			continue
		}

		var obj ServerConfig
		err := json.Unmarshal(jsonBody, &obj)
		if err == nil {
			server, err1 := NewServer(m.name, obj)
			if err1 == nil {
				m.servers[obj.Name] = server

				serverNames = append(serverNames, obj.Name)
			}
		}
	}

	defaultS, err := config.GetManager().Get(m.name, "default")
	if err == nil {
		if utils.ArrayContains(&serverNames, defaultS.(string)) {
			m.defaultServer = defaultS.(string)
		} else if len(serverNames) > 0 {
			m.defaultServer = serverNames[0]
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

// AddRoute - add a route to the server with specified name
func (m *manager) AddRoute(method string, path string, f func(c *fiber.Ctx) error, routeName string, versions []string, groupNames []string, serverName ...string) error {
	if len(serverName) > 0 {
		for _, sn := range serverName {
			if s, ok := m.servers[sn]; ok {
				return s.AddRoute(method, path, f, routeName, versions, groupNames)
			}
		}
	} else {
		if m.defaultServer != "" {
			return m.servers[m.defaultServer].AddRoute(method, path, f, routeName, versions, groupNames)
		}
	}
	return NewAddRouteToNilServerErr(path)
}

func (m *manager) GetRouteByName(routeName string, serverName ...string) (*fiber.Route, error) {
	if len(serverName) > 1 {
		return nil, NewFromMultipleServerErr()
	} else if len(serverName) == 1 {
		for _, sn := range serverName {
			if s, ok := m.servers[sn]; ok {
				return s.GetRouteByName(routeName)
			}
		}
	} else {
		if m.defaultServer != "" {
			return m.servers[m.defaultServer].GetRouteByName(routeName)
		}
	}
	return nil, NewFromNilServerErr()
}

// AddGroup - add a group to the server with specified name
func (m *manager) AddGroup(groupName string, f func(c *fiber.Ctx) error, groupsName []string, serverName ...string) error {
	if serverName != nil {
		if len(serverName) > 0 {
			for _, sn := range serverName {
				if s, ok := m.servers[sn]; ok {
					if groupsName != nil {
						return s.AddGroup(groupName, f, groupsName...)
					} else {
						return s.AddGroup(groupName, f, []string{}...)
					}
				}
			}
		} else {
			if m.defaultServer != "" {
				if groupsName != nil {
					return m.servers[m.defaultServer].AddGroup(groupName, f, groupsName...)
				} else {
					return m.servers[m.defaultServer].AddGroup(groupName, f, []string{}...)
				}
			}
		}
	} else {
		if m.defaultServer != "" {
			if groupsName != nil {
				return m.servers[m.defaultServer].AddGroup(groupName, f, groupsName...)
			} else {
				return m.servers[m.defaultServer].AddGroup(groupName, f, []string{}...)
			}
		}
	}
	return NewAddGroupToNilServerErr(groupName)
}

// AttachErrorHandler - attach a custom error handler to the server with specified name
func (m *manager) AttachErrorHandler(f func(ctx *fiber.Ctx, err error) error, serverNames ...string) error {
	if len(serverNames) > 0 {
		for _, sn := range serverNames {
			if s, ok := m.servers[sn]; ok {
				s.AttachErrorHandler(f)
				return nil
			}
		}
	} else {
		if m.defaultServer != "" {
			m.servers[m.defaultServer].AttachErrorHandler(f)
			return nil
		}
	}

	return NewAttachErrorHandlerToNilServerErr(serverNames...)
}
