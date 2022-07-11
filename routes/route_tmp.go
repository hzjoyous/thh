package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"thh/arms/logger"
)

func regTmpRoutes(r *gin.Engine) {
	tmp := r.Group("tmp")
	tmp.Any("log", func(context *gin.Context) {
		max := 10
		for i := 0; i < max; i++ {
			logger.Std().Info("log test")
		}
		context.String(200, "okokokooko")
	})
	tmp.Any("showRequest", func(c *gin.Context) {
		buf := make([]byte, 1024)
		n, _ := c.Request.Body.Read(buf)
		fmt.Println(string(buf[0:n]))
		c.JSON(http.StatusOK, map[string]any{"ok": "ok"})
	})
}
