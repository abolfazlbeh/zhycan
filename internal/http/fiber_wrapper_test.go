package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"testing"
)

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

func TestFiberWrapper_AddRoute(t *testing.T) {
	serverConfig := ServerConfig{ListenAddress: ":3000"}
	server, err := NewServer(serverConfig)
	if err != nil {
		t.Errorf("Creating HTTP Server --> Expected: %v, but got %v", nil, err)
	}

	server.AddRoute(fiber.MethodGet, "/", func(c *fiber.Ctx) error {
		return nil
	}, "TestGet")

	routes := server.app.GetRoutes()
	if len(routes) == 0 {
		t.Errorf("Expected 1 routes added to server, but got %v", len(routes))
	}

	// check the name of the route
	r := server.app.GetRoute("TestGet")
	if r.Method != fiber.MethodGet {
		t.Errorf("Expected to get method %v, but got %v", fiber.MethodGet, r.Method)
	}
	if r.Path != "/" {
		t.Errorf("Expected to get route '%v', but got '%v'", "/", r.Path)
	}
}

func TestFiberWrapper_AddMultipleRoute(t *testing.T) {
	serverConfig := ServerConfig{ListenAddress: ":3000"}
	server, err := NewServer(serverConfig)
	if err != nil {
		t.Errorf("Creating HTTP Server --> Expected: %v, but got %v", nil, err)
	}

	for i := range []int{1, 2, 3} {
		server.AddRoute(fiber.MethodGet, fmt.Sprintf("/%v", i), func(c *fiber.Ctx) error {
			return nil
		}, fmt.Sprintf("Test%v", i))
	}

	routes := server.app.GetRoutes()
	if len(routes) == 0 {
		t.Errorf("Expected 1 routes added to server, but got %v", len(routes))
	} else if len(routes) != 3 {
		t.Errorf("Expected 1 routes added to server, but got %v", len(routes))
	}

	// check the name of the route
	r := server.app.GetRoute("Test2")
	if r.Method != fiber.MethodGet {
		t.Errorf("Expected to get method %v, but got %v", fiber.MethodGet, r.Method)
	}
	if r.Path != "/2" {
		t.Errorf("Expected to get route '%v', but got '%v'", "/2", r.Path)
	}
}

func TestFiberWrapper_AddBlockRouteMethod(t *testing.T) {
	serverConfig := ServerConfig{ListenAddress: ":3000"}
	server, err := NewServer(serverConfig)
	if err != nil {
		t.Errorf("Creating HTTP Server --> Expected: %v, but got %v", nil, err)
	}

	server.AddRoute("TTT", "/", func(c *fiber.Ctx) error {
		return nil
	}, "TestGet")

	routes := server.app.GetRoutes()
	if len(routes) > 0 {
		t.Errorf("Expected the route not to be added, but added")
	}
}
