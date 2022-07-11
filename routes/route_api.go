package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thh/app/http/controllers"
	"thh/app/http/middleware"
)

func registerApiRoutes(r *gin.Engine) {
	apiGroup := r.Group("api")
	{
		apiGroup.GET("/", controllers.About)
		apiGroup.GET("/showPic", controllers.ShowPic)
		apiGroup.GET("/showPic2", controllers.ShowPic2)
		apiGroup.POST("/reg", controllers.Register)
		apiGroup.POST("/login", controllers.Login)
		apiGroup.GET("/setData", controllers.SetData)
		apiGroup.GET("/getData", controllers.GetData)
		apiGroup.POST("/getUserInfo", upHandle(controllers.GetUserInfo))
		apiGroup.POST("/upload", controllers.Upload)
		apiGroup.GET("memUse", upHandle(controllers.GetUseMem))
		apiGroup.GET("logTest", controllers.LoggerTest)
	}

	authApi := r.Group("api", middleware.JWTAuth())
	{
		authApi.GET("/userInfo", upHandleAuth(controllers.UserInfoV2))
	}
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
