package conf

import "thh/helpers/config"

const on = "on"
const off = "off"

func UseMigration() bool {
	openMigration := config.GetString("database.mysql.openMigration")
	return openMigration == on
}

func init() {
	config.Add("database", func() map[string]interface{} {

		return map[string]interface{}{
			// 默认使用的链接方式
			"default": config.Env("DB_CONNECTION", "sqlite"),

			"sqlite": map[string]interface{}{
				// sqlite 路径
				"path": config.Env("DB_PATH", "./storage/sqlite.DB"),
			},

			"mysql": map[string]interface{}{
				// 是否使用数据迁移
				"openMigration": config.Env("OPEN_MIGRATION", "off"),

				// mysql 配置地址
				"url": config.Env("DATABASE_URL", "db_user:db_pass@tcp(db_host:3306)/db_name?charset=utf8mb4&parseTime=True&loc=Local"),

				// 连接池配置
				"max_idle_connections": config.Env("DB_MAX_IDLE_CONNECTIONS", 20),

				// 最大连接数
				"max_open_connections": config.Env("DB_MAX_OPEN_CONNECTIONS", 20),

				// 最大生存时间
				"max_life_seconds":     config.Env("DB_MAX_LIFE_SECONDS", 300),
			},
		}
	})
}
