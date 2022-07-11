package routes

import (
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
	"path"
	"thh/app/http/controllers"
	"thh/arms/app"
)

type fsFunc func(name string) (fs.File, error)

func (f fsFunc) Open(name string) (fs.File, error) {
	return f(name)
}

func registerWebRoutes(r *gin.Engine) {

	// 获取打包的静态资源
	webFS := app.GetWebFS()

	// 设置静态资源路径
	r.StaticFS("/web", http.FS(webFS))

	actorFS := app.GetActorFS()

	// 设置线上静态资源路径
	handler := fsFunc(func(name string) (fs.File, error) {
		assetPath := path.Join("./actor/dist", name)
		// If we can't find the asset, fs can handle the error
		file, err := actorFS.Open(assetPath)
		if err != nil {
			return nil, err
		}
		// Otherwise, assume this is a legitimate request routed correctly
		return file, err
	})

	// 获取静态资源
	if app.IsProduction() {
		r.StaticFS("/actor", http.FS(handler))
	} else {
		r.Static("/actor", "./actor/dist")
	}

	r.GET("GetClashConfig", controllers.GetClashConfig)

}
