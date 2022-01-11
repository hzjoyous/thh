package bootstrap

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"thh/app/console"
	"thh/app/models/dataRep"
	"thh/app/models/user"
	"thh/conf"
	"thh/helpers/config"
	"thh/helpers/db"
	"thh/helpers/logger"
	"thh/routes"
	"time"
)

type Application struct {
}

var App Application

func Initialize() Application {

	// 预加载项目配置加载配置文件
	conf.Initialize()

	// 读取项目配置
	config.InitConfig("")

	// 日志
	Logger.Init(conf.LogPath())

	// 数据库链接迁移
	DB.InitConnection()

	// 数据库迁移
	migration(conf.UseMigration(), DB.SqlDBIns())

	// 初始化应用程序
	App = Application{}

	if config.GetBool("app.debug", true) {
		go func() {
			//http://127.0.0.1:6060/debug/pprof/
			err := http.ListenAndServe("0.0.0.0:6060", nil)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}

	return App
}

func migration(migration bool, db *gorm.DB) {
	if migration == false {
		return
	}
	// 自动迁移
	var err error

	if err = db.AutoMigrate(
		&user.User{},
		&dataRep.DataRep{},
	); err != nil {
		log.Println(err)
	}
}

// ServerRun 运行程序
func (itself *Application) ServerRun() {

	go console.RunJob()

	port := config.GetString("app.port")
	var engine *gin.Engine
	switch conf.IsProd() {
	case true:
		gin.DisableConsoleColor()
		gin.SetMode(gin.ReleaseMode)
		engine = gin.New()
		break
	default:
		engine = gin.Default()
		break
	}

	routes.RegisterRoutes(engine)

	srv := &http.Server{
		Addr:           ":" + port,
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Thousand-hand:listen " + port)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	// 这一句是 go chan 等待接受值，只是把接到的值直接扔掉了，此处是主协程的阻塞处
	_ = <-quit

	Logger.Std().Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		Logger.Std().Println("Server Shutdown:", err)
	}
	Logger.Std().Println("Server exiting")
}
