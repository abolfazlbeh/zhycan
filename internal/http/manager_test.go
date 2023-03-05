package http

import (
	"github.com/gofiber/fiber/v2"
	"reflect"
	"testing"
	"zhycan/internal/config"
)

func TestManager_Init(t *testing.T) {
	makeReadyConfigManager()

	m := manager{}
	m.init()

	if m.name != "http" {
		t.Errorf("Expected manager name to be 'http', got '%s'", m.name)
	}
	if len(m.servers) == 0 {
		t.Errorf("Expected manager to have at least one server, got '%d'", len(m.servers))
	}
}

func TestGetManager(t *testing.T) {
	m := GetManager()
	if m == nil {
		t.Errorf("Expected manager instance to be not nil, got '%v'", m)
	}
}

func TestManger_StartServers(t *testing.T) {
	makeReadyConfigManager()

	// Get instance of manager
	err := GetManager().StartServers()
	if err != nil {
		t.Errorf("Starting HTTP Manager --> Expected: %v, but got %v", nil, err)
	}
}

func TestManager_StopServers(t *testing.T) {
	makeReadyConfigManager()

	// Get instance of manager
	err := GetManager().StopServers()
	if err != nil {
		t.Errorf("Stopping HTTP Manager --> Expected: %v, but got %v", nil, err)
	}
}

func TestManager_CheckServerConfig(t *testing.T) {
	makeReadyConfigManager()

	// Get the first server
	m := GetManager()
	if len(m.servers) == 0 {
		t.Errorf("Expected manager to have at least one server, got '%d'", len(m.servers))
	}

	// check the configs
	expectedAddress := ":3000"
	if !reflect.DeepEqual(m.servers[m.defaultServer].config.ListenAddress, expectedAddress) {
		t.Errorf("Expected the Addr of the first server to be %v, but got %v", expectedAddress, m.servers[m.defaultServer].config.ListenAddress)
	}

	expectedVal := ".gz"
	if !reflect.DeepEqual(m.servers[m.defaultServer].config.Config.CompressedFileSuffix, expectedVal) {
		t.Errorf("Expected the config val of the first server to be %v, but got %v", expectedVal, m.servers[m.defaultServer].config.Config.CompressedFileSuffix)
	}

}

func TestManager_AddRoute(t *testing.T) {
	makeReadyConfigManager()

	// Get the first server
	m := GetManager()
	if len(m.servers) == 0 {
		t.Errorf("Expected manager to have at least one server, got '%d'", len(m.servers))
	}

	err := m.AddRoute(fiber.MethodGet, "/", func(c *fiber.Ctx) error {
		return nil
	}, "TestGet", []string{}, []string{})
	if err != nil {
		t.Errorf("Adding route to server expected to return %v, but got %v", nil, err)
	}

	routes := m.servers[m.defaultServer].app.GetRoutes()
	if len(routes) == 0 {
		t.Errorf("Expected 1 routes added to server, but got %v", len(routes))
	}
}

func TestManager_AddRouteToSpecificServer(t *testing.T) {
	makeReadyConfigManager()

	// Get the first server
	m := GetManager()
	if len(m.servers) == 0 {
		t.Errorf("Expected manager to have at least one server, got '%d'", len(m.servers))
	}

	err := m.AddRoute(fiber.MethodGet, "/", func(c *fiber.Ctx) error {
		return nil
	}, "TestGet", []string{}, []string{})
	if err != nil {
		t.Errorf("Adding route to server expected to return %v, but got %v", nil, err)
	}

	routes := m.servers["s1"].app.GetRoutes()
	if len(routes) == 0 {
		t.Errorf("Expected 1 routes added to server %v, but got %v", "s1", len(routes))
	}
}

func TestManager_GetRouteByName(t *testing.T) {
	makeReadyConfigManager()

	// Get the first server
	m := GetManager()
	if len(m.servers) == 0 {
		t.Errorf("Expected manager to have at least one server, got '%d'", len(m.servers))
		return
	}

	serverName := "s1"
	routeName := "TestGet"
	err := m.AddRoute(fiber.MethodGet, "/", func(c *fiber.Ctx) error {
		return nil
	}, routeName, []string{}, []string{}, serverName)
	if err != nil {
		t.Errorf("Adding route to server expected to return %v, but got %v", nil, err)
		return
	}

	route, err := m.GetRouteByName(routeName)
	if err != nil {
		t.Errorf("Error Get route of server by name, expected %v, but got %v", nil, err)
		return
	}

	if route.Name != routeName {
		t.Errorf("Expected to get the route name %v, but got %v", routeName, route.Name)
		return
	}

	route1, err := m.GetRouteByName(routeName, serverName)
	if err != nil {
		t.Errorf("Error Get route of server by name, expected %v, but got %v", nil, err)
		return
	}

	if route1.Name != routeName {
		t.Errorf("Expected to get the route name %v, but got %v", routeName, route1.Name)
		return
	}
}

func makeReadyConfigManager() {
	path := "../.."
	initialMode := "test"
	prefix := "ZHYCAN"

	_ = config.CreateManager(path, initialMode, prefix)
}
