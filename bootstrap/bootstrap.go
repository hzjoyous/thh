package bootstrap

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"thh/app/models/ActivityConfig"
	"thh/app/models/ActivityLimitConfig"
	"thh/app/models/Permission"
	"thh/app/models/Role"
	"thh/app/models/RolePermission"
	"thh/app/models/Users"
	"thh/app/models/dataRep"
	"thh/arms"
	"thh/arms/app"
	"thh/arms/config"
	"thh/arms/logger"
	"thh/conf"
	dbconnect2 "thh/conf/dbconnect"
)

func Initialize() {

	// 预加载项目配置加载配置文件
	conf.Initialize()
	if !arms.IsExist("./.env") {
		err := arms.Put([]byte(app.GetEnvExample()), "./.env")
		if err != nil {
			panic(err)
		}
	}
	// 读取项目配置
	config.InitConfig("")

	// 日志
	logger.Init(conf.LogPath())

	// 数据库链接
	dbconnect2.InitConnection()

	// 数据库迁移
	migration(conf.UseMigration(), dbconnect2.Std())

	arms.SetBasePath("storage/")
}

func migration(migration bool, db *gorm.DB) {
	if migration == false {
		return
	}
	// 自动迁移
	var err error

	if err = db.AutoMigrate(
		&dataRep.DataRep{},
		&Users.Users{},
		&Users.User{},
		&Role.Role{},
		&RolePermission.RolePermission{},
		&Permission.Permission{},
		&ActivityConfig.ActivityConfig{},
		&ActivityLimitConfig.ActivityLimitConfig{},
	); err != nil {
		log.Println(err)
	} else {
		fmt.Println("migration end")
	}
}
