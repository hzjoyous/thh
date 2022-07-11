package Permission

import (
	"gorm.io/gorm"
	db "thh/conf/dbconnect"
)

func builder() *gorm.DB {
	return db.Std().Table(tableName)
}

func First(db *gorm.DB) (el Permission) {
	db.First(&el)
	return
}

func List(db *gorm.DB) (el []Permission) {
	db.Find(&el)
	return
}
