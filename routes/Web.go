package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rakyll/statik/fs"
	"log"
	"thh/app/http/controllers"
	"thh/conf"
	_ "thh/statik"
)

func registerWebRoutes(r *gin.Engine) {

	if conf.IsProd() {
		statikFS, err := fs.New()
		if err != nil {
			log.Fatal(err)
		}
		// 静态文件
		r.StaticFS("/web", statikFS)
	} else {
		r.Static("/web", "./web")
	}
	r.Static("/actor", "./actor/dist")

	r.GET("GetClashConfig", controllers.GetClashConfig)

}
