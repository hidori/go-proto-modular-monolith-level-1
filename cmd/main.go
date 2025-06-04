package main

import (
	"context"

	"github.com/gin-gonic/gin"
	router1 "github.com/hidori/go-proto-modular-monolith-level-1/app1/public/infra/router"
	router2 "github.com/hidori/go-proto-modular-monolith-level-1/app2/public/infra/router"
)

func main() {
	ctx := context.Background()

	server := gin.Default()
	router1.Attach(ctx, server.Group("/app1"))
	router2.Attach(ctx, server.Group("/app2"))

	server.Run(":8080")
}
