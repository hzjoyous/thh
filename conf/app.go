package conf

import "thh/arms/config"

const (
	ENV      = "env"
	EnvProd  = "production"
	EnvLocal = "local"
	Port     = "port"

	LogTypeStdout = "stdout"
	LogTypeFile   = "file"
)

func init() {
	config.Add("app", func() map[string]any {
		return map[string]any{
			// 应用名称，暂时没有使用到
			"name": config.Env("APP_NAME", "THH"),

			// 当前环境，用以区分多环境
			"env": config.Env("APP_ENV", EnvProd),

			// 是否进入调试模式
			"debug": config.Env("APP_DEBUG", false),

			// 应用服务端口
			"port": config.Env("APP_PORT", "8080"),

			// 用以生成链接
			"url": config.Env("APP_URL", "http://localhost:8080"),

			// 日志输出方式，默认使用标准输出
			"logType": config.Env("LOG_TYPE", LogTypeStdout),

			// log 地址
			"logPath": config.Env("LOG_PATH", "./tmp/app.log"),

			// 加密会话、JWT 加密
			"key": config.Env("APP_KEY", "33446a9dcf9ea060a0a6532b166da32f304af0de"),

			// 设置时区，JWT 里会使用，日志记录里也会使用到
			"timezone": config.Env("TIMEZONE", "Asia/Shanghai"),

			// API 域名，未设置的话所有 API URL 加 api 前缀，如 http://domain.com/api/v1/users
			"api_domain": config.Env("API_DOMAIN"),
		}
	})
}

func IsProd() bool {
	return config.Get("app.env") == EnvProd
}

func LogPath() string {
	return config.Get("app.logPath")
}

func LogType() string {
	return config.Get("app.logType")
}
