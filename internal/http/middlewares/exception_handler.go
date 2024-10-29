package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
)

type ErrorHttpResponse struct {
	Message     string
	Status      int
	Description string
}

func ErrorHandlerMiddleware(c *gin.Context, err any) {
	goErr := errors.Wrap(err, 2)
	httpResponse := ErrorHttpResponse{
		Message:     "Internal server error",
		Status:      500,
		Description: goErr.Error(),
	}
	c.AbortWithStatusJSON(500, httpResponse)
}
