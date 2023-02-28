package shortcuts

import (
	"github.com/gofiber/fiber/v2"
	"zhycan/internal/http"
)

// HttpRoute - Structure of the route
type HttpRoute struct {
	Method    string
	Path      string
	RouteName string
	Versions  []string
	F         *func(c *fiber.Ctx) error
	Servers   []string
}

// AddHttpRouteByObj - add route by HttpRoute obj
func AddHttpRouteByObj(httpRoute HttpRoute) error {
	return http.GetManager().AddRoute(httpRoute.Method,
		httpRoute.Path,
		*httpRoute.F,
		httpRoute.RouteName,
		httpRoute.Versions,
		httpRoute.Servers...)
}

// AddHttpRoute - Add route by parameters
func AddHttpRoute(method string, path string, f func(c *fiber.Ctx) error, routeName string, versions []string, serverName ...string) error {
	return http.GetManager().AddRoute(method,
		path,
		f,
		routeName,
		versions,
		serverName...)
}

// AddBulkHttpRoutes - add bulk http routes to the server
func AddBulkHttpRoutes(httpRoutes []HttpRoute) error {
	for _, httpRoute := range httpRoutes {
		err := http.GetManager().AddRoute(httpRoute.Method,
			httpRoute.Path,
			*httpRoute.F,
			httpRoute.RouteName,
			httpRoute.Versions,
			httpRoute.Servers...)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetRouteByName(routeName string, serverName ...string) (*fiber.Route, error) {
	return http.GetManager().GetRouteByName(routeName, serverName...)
}
