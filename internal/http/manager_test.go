package http

import (
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
	if !reflect.DeepEqual(m.servers[0].config.ListenAddress, expectedAddress) {
		t.Errorf("Expected the Addr of the first server to be %v, but got %v", expectedAddress, m.servers[0].config.ListenAddress)
	}

	expectedVal := ".gz"
	if !reflect.DeepEqual(m.servers[0].config.Config.CompressedFileSuffix, expectedVal) {
		t.Errorf("Expected the config val of the first server to be %v, but got %v", expectedVal, m.servers[0].config.Config.CompressedFileSuffix)
	}

}

func makeReadyConfigManager() {
	path := "../.."
	initialMode := "test"
	prefix := "ZHYCAN"

	_ = config.CreateManager(path, initialMode, prefix)
}
