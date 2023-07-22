package engine

import (
	"github.com/abolfazlbeh/zhycan/internal/engine"
	"github.com/abolfazlbeh/zhycan/pkg/http"
	"strings"
)

// RegisterRestfulApp - register the restful application to the engine
func RegisterRestfulApp(app engine.RestfulApp) error {
	routes := app.Routes()

	controllerName := app.GetName()
	if strings.TrimSpace(controllerName) != "" {
		http.AddHttpGroupByObj(http.HttpGroup{
			GroupName: controllerName,
			F:         nil,
		})

		for i, item := range routes {
			routes[i].GroupNames = append([]string{controllerName}, item.GroupNames...)
		}
	}

	http.AddBulkHttpRoutes(routes)
	return nil
}

// Controller Struct - The controller struct to be used by others
type Controller struct {
	name string
}

func (c Controller) GetName() string {
	return c.name
}
