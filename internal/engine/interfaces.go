package engine

import "github.com/abolfazlbeh/zhycan/pkg/http"

type RestfulApp interface {
	Routes() []http.HttpRoute
	GetName() string
}
