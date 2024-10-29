package middlewares

import (
	"github.com/abolfazlbeh/zhycan/internal/http/types"
	"github.com/fufuok/favicon"
	"github.com/gin-gonic/gin"
)

func FaviconMiddleware(config types.FaviconMiddlewareConfig) gin.HandlerFunc {
	return favicon.New(favicon.Config{
		File:         config.File,
		CacheControl: config.CacheControl,
	})
}
