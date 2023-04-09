package http

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
	"strings"
	"time"
	"zhycan/internal/config"
	"zhycan/internal/utils"
)

// Mark: Definitions

// Server struct
type Server struct {
	name                  string
	config                ServerConfig
	app                   *fiber.App
	versionGroups         map[string]fiber.Router
	groups                map[string]fiber.Router
	supportedMiddlewares  []string
	defaultRequestMethods []string
}

// init - Server Constructor - It initializes the server
func (s *Server) init(name string, serverConfig ServerConfig) error {
	s.name = name
	s.config = serverConfig

	// Get application name from the config manager
	appName := config.GetManager().GetName()
	requestMethods := fiber.DefaultMethods
	if s.config.Config.RequestMethods[0] != "ALL" {
		requestMethods = s.config.Config.RequestMethods
	}
	s.defaultRequestMethods = requestMethods

	s.app = fiber.New(fiber.Config{
		Prefork:        false,
		AppName:        appName,
		RequestMethods: requestMethods,
	})
	s.groups = make(map[string]fiber.Router)
	s.supportedMiddlewares = []string{
		"logger",
		"favicon",
	}
	return nil
}

func (s *Server) createVersionGroups(versions []string) {
	s.versionGroups = make(map[string]fiber.Router)
	for _, item := range versions {
		s.versionGroups[item] = s.app.Group(item)
	}
}

func (s *Server) attachMiddlewares(orders []string) {
	for _, item := range orders {
		if utils.ArrayContains(&s.supportedMiddlewares, item) {
			switch item {
			case "logger":
				key := fmt.Sprintf("middlewares.%s", item)
				// read config
				loggerCfg, err := config.GetManager().Get(s.name, key)
				if err == nil {
					jsonBody, err2 := json.Marshal(loggerCfg.(interface{}))
					if err2 == nil {
						var obj LoggerMiddlewareConfig
						err := json.Unmarshal(jsonBody, &obj)
						if err == nil {
							// Everything is ok and let's go define logger config
							loggerMiddlewareCfg := logger.Config{
								Next:         nil,
								Done:         nil,
								Format:       obj.Format,
								TimeFormat:   obj.TimeFormat,
								TimeZone:     obj.TimeZone,
								TimeInterval: time.Duration(obj.TimeInterval) * time.Millisecond,
							}
							if obj.Output == "stdout" {
								loggerMiddlewareCfg.Output = os.Stdout
							}
							s.app.Use(logger.New(loggerMiddlewareCfg))
							break
						}
					}
				}
				s.app.Use(logger.New())

			case "favicon":
				key := fmt.Sprintf("middlewares.%s", item)
				// read config
				loggerCfg, err := config.GetManager().Get(s.name, key)
				if err == nil {
					jsonBody, err2 := json.Marshal(loggerCfg.(interface{}))
					if err2 == nil {
						var obj FaviconMiddlewareConfig
						err := json.Unmarshal(jsonBody, &obj)
						if err != nil {
							faviconMiddlewareCfg := favicon.Config{
								File:         obj.File,
								URL:          obj.URL,
								CacheControl: obj.CacheControl,
							}
							s.app.Use(favicon.New(faviconMiddlewareCfg))
							break
						}
					}
				}
				s.app.Use(favicon.New())
			}
		}
	}
}

func (s *Server) addGroup(keyName string, groupName string, router fiber.Router, f func(c *fiber.Ctx) error) {
	if f == nil {
		s.groups[keyName] = router.Group(groupName)
	} else {
		s.groups[keyName] = router.Group(groupName, f)
	}
}

// MARK: Public functions

// NewServer - create a new instance of Server and return it
func NewServer(name string, config ServerConfig) (*Server, error) {
	server := &Server{}
	err := server.init(name, config)
	if err != nil {
		return nil, NewCreateServerErr(err)
	}

	server.attachMiddlewares(config.Middlewares.Order)
	server.createVersionGroups(config.Versions)
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
func (s *Server) AddRoute(method string, path string, f func(c *fiber.Ctx) error, routeName string, versions []string, groups []string) error {
	// check that whether is acceptable to add this route method
	if utils.ArrayContains(&s.defaultRequestMethods, method) {
		if len(groups) > 0 {
			for _, g := range groups {
				if len(versions) > 0 {
					for _, v := range versions {
						if v == "all" {
							for k := range s.versionGroups {
								newKey := fmt.Sprintf("%s.%s", k, g)
								if router, ok := s.groups[newKey]; ok {
									router.Add(method, path, f)
									if strings.TrimSpace(routeName) != "" {
										//router.Name(routeName)
										s.app.Name(routeName)
									}
								}
							}
							break
						} else if v == "" {
							if router, ok := s.groups[g]; ok {
								router.Add(method, path, f)
								if strings.TrimSpace(routeName) != "" {
									//router.Name(routeName)
									s.app.Name(routeName)
								}
							}
							break
						} else {
							newKey := fmt.Sprintf("%s.%s", v, g)
							if router, ok := s.groups[newKey]; ok {
								router.Add(method, path, f)
								if strings.TrimSpace(routeName) != "" {
									s.app.Name(routeName)
									//router.Name(routeName)
								}
							}
						}
					}
				} else {
					if savedGroup, ok := s.groups[g]; ok {
						savedGroup.Add(method, path, f)
						if strings.TrimSpace(routeName) != "" {
							//savedGroup.Name(routeName)
							s.app.Name(routeName)

						}
					}
				}
			}
		} else {
			if len(versions) > 0 {
				for _, v := range versions {
					if router, ok := s.versionGroups[v]; ok {
						router.Add(method, path, f)
						if strings.TrimSpace(routeName) != "" {
							//router.Name(routeName)
							s.app.Name(routeName)
						}
					} else {
						if v == "all" {
							for _, router1 := range s.versionGroups {
								router1.Add(method, path, f)
								if strings.TrimSpace(routeName) != "" {
									//router1.Name(routeName)
									s.app.Name(routeName)
								}
							}
							break
						} else if v == "" {
							s.app.Add(method, path, f)
							if strings.TrimSpace(routeName) != "" {
								s.app.Name(routeName)
							}
							break
						}
					}
				}
			} else {
				s.app.Add(method, path, f)
				if strings.TrimSpace(routeName) != "" {
					s.app.Name(routeName)
				}
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
			for key := range s.versionGroups {
				gKey := fmt.Sprintf("%s.%s", key, g)
				if r, ok := s.groups[gKey]; ok {
					newKey := fmt.Sprintf("%s.%s.%s", key, g, groupName)
					s.addGroup(newKey, groupName, r, f)
				} else {
					return NewGroupRouteNotExistErr(gKey)
				}
			}

			newKey := fmt.Sprintf("%s.%s", g, groupName)
			s.addGroup(newKey, groupName, s.app, f)
		}
	} else {
		for key, item := range s.versionGroups {
			newKey := fmt.Sprintf("%s.%s", key, groupName)
			s.addGroup(newKey, groupName, item, f)
		}

		s.addGroup(groupName, groupName, s.app, f)
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
