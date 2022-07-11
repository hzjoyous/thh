package ActivityConfig

import (
	"gorm.io/gorm"
	db "thh/conf/dbconnect"
)

func builder() *gorm.DB {
	return db.Std().Table(tableName)
}

func First(db *gorm.DB) (el ActivityConfig) {
	db.First(&el)
	return
}

func List(db *gorm.DB) (el []ActivityConfig) {
	db.Find(&el)
	return
}
