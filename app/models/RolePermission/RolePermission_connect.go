package RolePermission

import (
	"gorm.io/gorm"
	db "thh/conf/dbconnect"
)

func builder() *gorm.DB {
	return db.Std().Table(tableName)
}

func First(db *gorm.DB) (el RolePermission) {
	db.First(&el)
	return
}

func List(db *gorm.DB) (el []RolePermission) {
	db.Find(&el)
	return
}
