package bootstrap

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"thh/app/models/dataRep"
	"thh/app/models/guest"
	"thh/app/models/user"
	"thh/conf"
	"thh/helpers/config"
	"thh/helpers/db"
	"thh/helpers/logger"
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
	logger.Init(conf.LogPath())

	// 数据库链接迁移
	db.InitConnection()

	// 数据库迁移
	migration(conf.UseMigration(), db.SqlDBIns())

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
		&guest.Guest{},
	); err != nil {
		log.Println(err)
	}
}