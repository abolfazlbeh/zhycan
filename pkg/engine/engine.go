package engine

import (
	"github.com/abolfazlbeh/zhycan/internal/engine"
	interalgRPC "github.com/abolfazlbeh/zhycan/internal/grpc"
	"github.com/abolfazlbeh/zhycan/pkg/http"
	"google.golang.org/grpc"
	"strings"
)

// RegisterRestfulController - register the restful controller to the engine
func RegisterRestfulController(app engine.RestfulApp) {
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
}

// Controller Struct - The controller struct to be used by others
type Controller struct {
	name string
}

func (c Controller) GetName() string {
	return c.name
}

func RegisterGrpcController(app engine.GrpcApp, registerFunc func(server *grpc.Server, cls interface{})) {
	controllerName := app.GetName()
	serverNames := app.GetServerNames()

	if strings.TrimSpace(controllerName) != "" {
		return
	}

	if len(serverNames) <= 0 {
		return
	}

	for _, item := range serverNames {
		srv, err := interalgRPC.GetManager().GetServerByName(item)
		if err != nil {
			// TODO: log the error
			continue
		}
		err = srv.RegisterController(registerFunc, app)
		if err != nil {
			// TODO: log the error
			continue
		}
	}
}
