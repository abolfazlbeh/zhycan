package db

import (
	"github.com/abolfazlbeh/zhycan/internal/config"
	"testing"
)

func TestManager_Init(t *testing.T) {
	makeReadyConfigManager()

	m := manager{}
	m.init()

	if m.name != "db" {
		t.Errorf("Expected manager name to be 'db', got '%s'", m.name)
	}
}

func makeReadyConfigManager() {
	path := "../.."
	initialMode := "test"
	prefix := "ZHYCAN"

	_ = config.CreateManager(path, initialMode, prefix)
}
