package http

import "testing"

func TestFiberWrapper_startServer(t *testing.T) {
	// create a new fiber wrapper and start it

	serverConfig := ServerConfig{ListenAddress: ":3000"}
	server, err := NewServer(serverConfig)
	if err != nil {
		t.Errorf("Creating HTTP Server --> Expected: %v, but got %v", nil, err)
	}

	err = server.Start()
	if err != nil {
		t.Errorf("Starting HTTP Server --> Expected: %v, but got %v", nil, err)
	}
}
