package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"reflect"
	"testing"
)

func TestFiberWrapper_startServer(t *testing.T) {
	// create a new fiber wrapper and start it

	serverConfig := ServerConfig{ListenAddress: ":3000"}
	server, err := NewServer(serverConfig)
	if err != nil {
		t.Errorf("Creating HTTP Server --> Expected: %v, but got %v", nil, err)
		return
	}

	err = server.Start()
	if err != nil {
		t.Errorf("Starting HTTP Server --> Expected: %v, but got %v", nil, err)
		return
	}
}

func TestFiberWrapper_AddRoute(t *testing.T) {
	serverConfig := ServerConfig{ListenAddress: ":3000"}
	server, err := NewServer(serverConfig)
	if err != nil {
		t.Errorf("Creating HTTP Server --> Expected: %v, but got %v", nil, err)
		return
	}

	server.AddRoute(fiber.MethodGet, "/", func(c *fiber.Ctx) error {
		return nil
	}, "TestGet", []string{})

	routes := server.app.GetRoutes()
	if len(routes) == 0 {
		t.Errorf("Expected 1 routes added to server, but got %v", len(routes))
		return
	}

	// check the name of the route
	r := server.app.GetRoute("TestGet")
	if r.Method != fiber.MethodGet {
		t.Errorf("Expected to get method %v, but got %v", fiber.MethodGet, r.Method)
		return
	}
	if r.Path != "/" {
		t.Errorf("Expected to get route '%v', but got '%v'", "/", r.Path)
		return
	}
}

func TestFiberWrapper_AddMultipleRoute(t *testing.T) {
	serverConfig := ServerConfig{ListenAddress: ":3000"}
	server, err := NewServer(serverConfig)
	if err != nil {
		t.Errorf("Creating HTTP Server --> Expected: %v, but got %v", nil, err)
		return
	}

	for i := range []int{1, 2, 3} {
		server.AddRoute(fiber.MethodGet, fmt.Sprintf("/%v", i), func(c *fiber.Ctx) error {
			return nil
		}, fmt.Sprintf("Test%v", i), []string{})
	}

	routes := server.app.GetRoutes()
	if len(routes) == 0 {
		t.Errorf("Expected 1 routes added to server, but got %v", len(routes))
		return
	} else if len(routes) != 3 {
		t.Errorf("Expected 1 routes added to server, but got %v", len(routes))
		return
	}

	// check the name of the route
	r := server.app.GetRoute("Test2")
	if r.Method != fiber.MethodGet {
		t.Errorf("Expected to get method %v, but got %v", fiber.MethodGet, r.Method)
		return
	}
	if r.Path != "/2" {
		t.Errorf("Expected to get route '%v', but got '%v'", "/2", r.Path)
		return
	}
}

func TestFiberWrapper_AddBlockRouteMethod(t *testing.T) {
	serverConfig := ServerConfig{ListenAddress: ":3000"}
	server, err := NewServer(serverConfig)
	if err != nil {
		t.Errorf("Creating HTTP Server --> Expected: %v, but got %v", nil, err)
		return
	}

	server.AddRoute("TTT", "/", func(c *fiber.Ctx) error {
		return nil
	}, "TestGet", []string{})

	routes := server.app.GetRoutes()
	if len(routes) > 0 {
		t.Errorf("Expected the route not to be added, but added")
		return
	}
}

func TestFiberWrapper_GetRouteByName(t *testing.T) {
	serverConfig := ServerConfig{ListenAddress: ":3000"}
	server, err := NewServer(serverConfig)
	if err != nil {
		t.Errorf("Creating HTTP Server --> Expected: %v, but got %v", nil, err)
		return
	}

	routeName := "TestGet"
	server.AddRoute(fiber.MethodGet, "/", func(c *fiber.Ctx) error {
		return nil
	}, routeName, []string{})

	routes := server.app.GetRoutes()
	if len(routes) == 0 {
		t.Errorf("Expected that %v route(s) exist, but got %v", 1, len(routes))
		return
	}

	route, err := server.GetRouteByName(routeName)
	routes = server.app.GetRoutes()
	if err != nil {
		t.Errorf("Get route by name --> Expected no err:  but got %v", err)
		return
	}

	if route.Name != routeName {
		t.Errorf("Expected to get route by name: %v, but got %v", routeName, route)
	}
}

func TestFiberWrapper_AddRouteToOneSupportedVersion(t *testing.T) {
	serverConfig := ServerConfig{ListenAddress: ":3000", Versions: []string{"v1"}}
	server, err := NewServer(serverConfig)
	if err != nil {
		t.Errorf("Creating HTTP Server --> Expected: %v, but got %v", nil, err)
		return
	}

	routeName := "TestGet"
	server.AddRoute(fiber.MethodGet, "/", func(c *fiber.Ctx) error {
		return nil
	}, routeName, []string{"v1"})

	routes := server.app.GetRoutes()
	if len(routes) == 0 {
		t.Errorf("Expected that %v route(s) exist, but got %v", 1, len(routes))
		return
	}

	expectedPath := "/v1"
	if routes[0].Path != expectedPath {
		t.Errorf("Expected that get route '%v', but got %v", expectedPath, routes[0].Path)
		return
	}
}

func TestFiberWrapper_AddRouteToTwoSupportedVersions(t *testing.T) {
	serverConfig := ServerConfig{ListenAddress: ":3000", Versions: []string{"v1", "v2"}}
	server, err := NewServer(serverConfig)
	if err != nil {
		t.Errorf("Creating HTTP Server --> Expected: %v, but got %v", nil, err)
		return
	}

	routeName := "TestGet"
	server.AddRoute(fiber.MethodGet, "/", func(c *fiber.Ctx) error {
		return nil
	}, routeName, []string{"v1", "v2"})

	routes := server.app.GetRoutes()
	if len(routes) != 2 {
		t.Errorf("Expected that %v route(s) exist, but got %v", 2, len(routes))
		return
	}

	expectedPath := []string{"/v1/", "/v2/"}
	var actualPath []string
	for _, r := range routes {
		actualPath = append(actualPath, r.Path)
	}
	if !reflect.DeepEqual(actualPath, expectedPath) {
		t.Errorf("Expected that get route '%v', but got %v", expectedPath, actualPath)
		return
	}
}

func TestFiberWrapper_AddRouteToAllVersions(t *testing.T) {
	serverConfig := ServerConfig{ListenAddress: ":3000", Versions: []string{"v1", "v2"}}
	server, err := NewServer(serverConfig)
	if err != nil {
		t.Errorf("Creating HTTP Server --> Expected: %v, but got %v", nil, err)
		return
	}

	routeName := "TestGet"
	server.AddRoute(fiber.MethodGet, "/", func(c *fiber.Ctx) error {
		return nil
	}, routeName, []string{"all"})

	routes := server.app.GetRoutes()
	if len(routes) != 2 {
		t.Errorf("Expected that %v route(s) exist, but got %v", 2, len(routes))
		return
	}

	expectedPath := []string{"/v1/", "/v2/"}
	var actualPath []string
	for _, r := range routes {
		actualPath = append(actualPath, r.Path)
	}
	if !reflect.DeepEqual(actualPath, expectedPath) {
		t.Errorf("Expected that get route '%v', but got %v", expectedPath, actualPath)
		return
	}
}

func TestFiberWrapper_AddRouteToNoVersions(t *testing.T) {
	serverConfig := ServerConfig{ListenAddress: ":3000", Versions: []string{"v1", "v2"}}
	server, err := NewServer(serverConfig)
	if err != nil {
		t.Errorf("Creating HTTP Server --> Expected: %v, but got %v", nil, err)
		return
	}

	routeName := "TestGet"
	server.AddRoute(fiber.MethodGet, "/", func(c *fiber.Ctx) error {
		return nil
	}, routeName, []string{""})

	routes := server.app.GetRoutes()
	if len(routes) != 1 {
		t.Errorf("Expected that %v route(s) exist, but got %v", 2, len(routes))
		return
	}

	expectedPath := []string{"/"}
	var actualPath []string
	for _, r := range routes {
		actualPath = append(actualPath, r.Path)
	}
	if !reflect.DeepEqual(actualPath, expectedPath) {
		t.Errorf("Expected that get route '%v', but got %v", expectedPath, actualPath)
		return
	}
}

func TestFiberWrapper_AddGroup(t *testing.T) {
	serverConfig := ServerConfig{ListenAddress: ":3000", Versions: []string{"v1", "v2"}}
	server, err := NewServer(serverConfig)
	if err != nil {
		t.Errorf("Creating HTTP Server --> Expected: %v, but got %v", nil, err)
		return
	}

	groupName := "p"
	server.AddGroup(groupName, func(c *fiber.Ctx) error {
		return nil
	})

	expectedGroups := []string{"v1.p", "v2.p", "p"}
	var actualGroups []string
	for key, _ := range server.groups {
		actualGroups = append(actualGroups, key)
	}

	if !reflect.DeepEqual(expectedGroups, actualGroups) {
		t.Errorf("Expected that get group list '%v', but got %v", expectedGroups, actualGroups)
		return
	}
}

func TestFiberWrapper_AddGroupToGroup(t *testing.T) {
	serverConfig := ServerConfig{ListenAddress: ":3000", Versions: []string{"v1", "v2"}}
	server, err := NewServer(serverConfig)
	if err != nil {
		t.Errorf("Creating HTTP Server --> Expected: %v, but got %v", nil, err)
		return
	}

	groupName := "p2"
	server.AddGroup(groupName, func(c *fiber.Ctx) error {
		return nil
	})

	groupName = "p"
	server.AddGroup(groupName, func(c *fiber.Ctx) error {
		return nil
	}, "p2")

	expectedGroups := []string{"p2.p", "v1.p2", "v2.p2", "p2", "v1.p2.p", "v2.p2.p"}
	var actualGroups []string
	for key, _ := range server.groups {
		actualGroups = append(actualGroups, key)
	}

	if !reflect.DeepEqual(expectedGroups, actualGroups) {
		t.Errorf("Expected that get group list '%v', but got %v", expectedGroups, actualGroups)
		return
	}
}
