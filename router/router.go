package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/13sai/gin-frame/controllers"
)

func Load(g *gin.Engine) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	g.GET("/health", controllers.Health)
	g.POST("/addScore", controllers.AddIntegral)
	g.GET("/getScore/:accountId", controllers.GetIntegral)

	return g
}
