package Users

import (
	"gorm.io/gorm"
	"thh/conf/dbconnect"
)

func builder() *gorm.DB {
	return dbconnect.Std().Table(tableName)
}

func First(db *gorm.DB) (el Users) {
	db.First(&el)
	return
}

func List(db *gorm.DB) (el []Users) {
	db.Find(&el)
	return
}
