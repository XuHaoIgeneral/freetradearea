package router

import (
	"github.com/gin-gonic/gin"
	"nodeserver/handler/sd"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.GET("/node/status", sd.HealthCheck)
	return g
}