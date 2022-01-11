package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thh/app/http/middleware"
)

func RegisterRoutes(r *gin.Engine) {

	r.Use(middleware.Cors())
	r.Use(middleware.LoggerMid())

	registerWebRoutes(r)
	registerApiRoutes(r)
	websocketRoute(r)

	r.NoRoute(func(c *gin.Context) {
		//c.Request.URL.Path = "/actor"
		//r.HandleContext(c)
		c.String(http.StatusNotFound, "Not router")
	})
}
