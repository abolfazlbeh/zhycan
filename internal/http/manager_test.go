package http

import (
	"testing"
	"zhycan/internal/config"
)

func Test_StartManager(t *testing.T) {
	makeReadyConfigManager()

	// Get instance of manager
	err := GetManager().StartServers()
	if err != nil {
		t.Errorf("Starting HTTP Manager --> Expected: %v, but got %v", nil, err)
	}
}

func makeReadyConfigManager() {
	path := "../.."
	initialMode := "test"
	prefix := "ZHYCAN"

	_ = config.CreateManager(path, initialMode, prefix)
}
