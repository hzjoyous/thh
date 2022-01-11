package DB

import (
	"github.com/syndtr/goleveldb/leveldb"
	"thh/helpers"
	"thh/helpers/config"
	"thh/helpers/logger"
)

type KVDBInterface interface {
	Get(string) string
	Set(string, string)
	Delete(string)
}

type kvDB struct {
	initSuccess    bool
	leveldbConnect *leveldb.DB
}

func (db *kvDB) Get(key string) string {
	if !db.initSuccess {
		return ""
	}
	v, _ := db.leveldbConnect.Get([]byte(key), nil)
	return helpers.ToString(v)
}

func (db *kvDB) Set(key string, value string) {
	if !(db.initSuccess) {
		return
	}
	_ = db.leveldbConnect.Put([]byte(key), []byte(value), nil)
}

func (db *kvDB) Delete(key string) {
	if !(db.initSuccess) {
		return
	}
	_ = db.leveldbConnect.Delete([]byte(key), nil)
}

var kvManager *kvDB

func connectLeveldb() {
	kvManager = &kvDB{}
	path := config.GetString("leveldb.path")
	leveldbEntity, err := leveldb.OpenFile(path, nil)
	if err != nil {
		Logger.Std().Error(err)
	} else {
		Logger.Std().Info("leveldb 初始化成功")
		kvManager.leveldbConnect = leveldbEntity
		kvManager.initSuccess = true
	}

}

func KVDB() KVDBInterface {
	return kvManager
}
