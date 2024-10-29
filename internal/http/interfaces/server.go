package interfaces

import (
	"context"
)

type IServer interface {
	Start() error
	Stop() error
	AddRoute(method string, path string, f func(c context.Context) error, routeName string, versions []string, groups []string) error
}
