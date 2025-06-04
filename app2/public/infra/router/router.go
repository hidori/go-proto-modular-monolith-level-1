package router

import (
	"context"

	"github.com/gin-gonic/gin"
)

func Attach(_ context.Context, rg *gin.RouterGroup) error {
	rg.GET("/", func(c *gin.Context) {
		c.String(200, "Hello from app2.\r\n")
	})

	return nil
}
