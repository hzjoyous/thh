package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"thh/app/http/middleware"
)

func RegisterRoutes(r *gin.Engine) {

	r.Use(middleware.Cors())
	r.Use(middleware.Logger())

	registerWebRoutes(r)
	registerApiRoutes(r)
	websocketRoute(r)

	r.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.Request.URL.Path = "/actor"
			r.HandleContext(c)
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
		//c.String(http.StatusNotFound, "Not router")
	})
}
