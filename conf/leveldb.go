package conf

import "thh/arms/config"

func init() {
	config.Add("leveldb", func() map[string]any {
		return map[string]any{
			// leveldb 路径
			"path": config.Env("LEVELDB_PATH", "./storage/leveldb"),
		}
	})
}
