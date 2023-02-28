package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strings"
	"zhycan/internal/utils"
)

// Mark: Definitions

// Server struct
type Server struct {
	config        ServerConfig
	app           *fiber.App
	versionGroups map[string]fiber.Router
	groups        map[string]fiber.Router
}

// init - Server Constructor - It initializes the server
func (s *Server) init(config ServerConfig) error {
	s.config = config
	s.app = fiber.New(fiber.Config{
		Prefork: false,
	})
	s.groups = make(map[string]fiber.Router)
	return nil
}

func (s *Server) createVersionGroups(versions []string) {
	s.versionGroups = make(map[string]fiber.Router)
	for _, item := range versions {
		s.versionGroups[item] = s.app.Group(item)
	}
}

func (s *Server) addGroup(groupName string, router fiber.Router, f func(c *fiber.Ctx) error) {
	if f == nil {
		s.groups[groupName] = router.Group(groupName)
	} else {
		s.groups[groupName] = router.Group(groupName, f)
	}
}

// MARK: Public functions

// NewServer - create a new instance of Server and return it
func NewServer(config ServerConfig) (*Server, error) {
	server := &Server{}
	err := server.init(config)
	server.createVersionGroups(config.Versions)
	if err != nil {
		return nil, NewCreateServerErr(err)
	}
	return server, nil
}

// Start - start the server and listen to provided address
func (s *Server) Start() error {
	err := s.app.Listen(s.config.ListenAddress)
	if err != nil {
		return NewStartServerErr(s.config.ListenAddress, err)
	}
	return nil
}

// Stop - stop the server
func (s *Server) Stop() error {
	err := s.app.Shutdown()
	if err != nil {
		return NewShutdownServerErr(err)
	}
	return nil
}

// AddRoute - add a route to the server
func (s *Server) AddRoute(method string, path string, f func(c *fiber.Ctx) error, routeName string, versions []string) error {
	// check that whether is acceptable to add this route method
	if utils.ArrayContains(&fiber.DefaultMethods, method) {
		if len(versions) > 0 {
			for _, v := range versions {
				if router, ok := s.versionGroups[v]; ok {
					router.Add(method, path, f)
					if strings.TrimSpace(routeName) != "" {
						router.Name(routeName)
					}
				} else {
					if v == "all" {
						for _, router1 := range s.versionGroups {
							router1.Add(method, path, f)
							if strings.TrimSpace(routeName) != "" {
								router1.Name(routeName)
							}
						}
					} else if v == "" {
						s.app.Add(method, path, f)
						if strings.TrimSpace(routeName) != "" {
							s.app.Name(routeName)
						}
					}
				}
			}
		} else {
			s.app.Add(method, path, f)
			if strings.TrimSpace(routeName) != "" {
				s.app.Name(routeName)
			}
		}
		return nil
	}

	return NewNotSupportedHttpMethodErr(method)
}

// AddGroup - add a group to the server
func (s *Server) AddGroup(groupName string, f func(c *fiber.Ctx) error, groups ...string) error {
	if len(groups) > 0 {
		for _, g := range groups {
			for key, _ := range s.versionGroups {
				gKey := fmt.Sprintf("%s.%s", key, g)
				if r, ok := s.groups[gKey]; ok {
					newKey := fmt.Sprintf("%s.%s.%s", key, g, groupName)
					s.addGroup(newKey, r, f)
				} else {
					return NewGroupRouteNotExistErr(gKey)
				}
			}

			newKey := fmt.Sprintf("%s.%s", g, groupName)
			s.addGroup(newKey, s.app, f)
		}
	} else {
		for key, item := range s.versionGroups {
			newKey := fmt.Sprintf("%s.%s", key, groupName)
			s.addGroup(newKey, item, f)
		}

		s.addGroup(groupName, s.app, f)
	}

	return nil
}

func (s *Server) GetRouteByName(name string) (*fiber.Route, error) {
	route := s.app.GetRoute(name)
	if route.Name != name {
		return nil, NewGetRouteByNameErr(name)
	}
	return &route, nil
}
