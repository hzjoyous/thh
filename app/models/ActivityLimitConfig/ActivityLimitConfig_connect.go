package ActivityLimitConfig

import (
	"gorm.io/gorm"
	db "thh/conf/dbconnect"
)

func builder() *gorm.DB {
	return db.Std().Table(tableName)
}

func First(db *gorm.DB) (el ActivityLimitConfig) {
	db.First(&el)
	return
}

func List(db *gorm.DB) (el []ActivityLimitConfig) {
	db.Find(&el)
	return
}
