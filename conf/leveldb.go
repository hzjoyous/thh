package conf

import "thh/helpers/config"

func init() {
	config.Add("leveldb", func() map[string]interface{} {
		return map[string]interface{}{
			// leveldb 路径
			"path": config.Env("LEVELDB_PATH", "./storage/leveldb"),
		}
	})
}
