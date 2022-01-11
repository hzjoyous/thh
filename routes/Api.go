package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"thh/app/http/controllers"
	"thh/app/http/middleware"
	Logger "thh/helpers/logger"
)

func registerApiRoutes(r *gin.Engine) {
	r.Group("api").
		GET("/", controllers.About).
		GET("/showPic", controllers.ShowPic).
		GET("/showPic2", controllers.ShowPic2).
		POST("/reg", controllers.Register).
		POST("/login", controllers.Login).
		GET("/setData", controllers.SetData).
		GET("/getData", controllers.GetData).
		POST("/getUserInfo", upHandle(controllers.GetUserInfo)).
		POST("/upload", controllers.Upload).
		Any("memUse", upHandle(controllers.GetUseMem)).
		Any("logTest", controllers.LoggerTest).
		Any("log", func(context *gin.Context) {
			max := 10
			for i := 0; i < max; i++ {
				Logger.Std().Info("log test")
			}
			context.String(200, "okokokooko")
		})

	r.Group("api").Use(middleware.JWTAuth()).
		Any("/userInfo", upHandleAuth(controllers.UserInfoV2))

	r.Group("api").Any("showRequest", func(c *gin.Context) {
		buf := make([]byte, 1024)
		n, _ := c.Request.Body.Read(buf)
		fmt.Println(string(buf[0:n]))
		c.JSON(http.StatusOK, map[string]interface{}{"ok": "ok"})
	})
	r.Group("api").Any("maraiManager", controllers.MaraiManager)
}

func upHandle(action func(request controllers.Request) controllers.Response) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := controllers.Request{}
		request.Context = c
		response := action(request)
		c.JSON(response.Code, response.Data)
	}
}

func upHandleAuth(action func(request controllers.Request) controllers.Response) func(c *gin.Context) {
	return func(c *gin.Context) {
		userIdFromContext, ok := c.Get("userId")
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"message": "un Login",
			})
			return
		}
		userId, ok := userIdFromContext.(uint64)
		request := controllers.Request{}
		request.Context = c
		request.UserId = userId
		response := action(request)
		c.JSON(response.Code, response.Data)
	}
}
