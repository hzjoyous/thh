package Role

import (
	"gorm.io/gorm"
	db "thh/conf/dbconnect"
)

func builder() *gorm.DB {
	return db.Std().Table(tableName)
}

func First(db *gorm.DB) (el Role) {
	db.First(&el)
	return
}

func List(db *gorm.DB) (el []Role) {
	db.Find(&el)
	return
}
