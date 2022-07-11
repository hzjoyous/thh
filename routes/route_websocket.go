package routes

import (
	"github.com/gin-gonic/gin"
	"thh/app/http/controllers"
)

func websocketRoute(r *gin.Engine) {
	r.GET("/ws", controllers.WsHandle)
	go controllers.Broadcaster()
}
