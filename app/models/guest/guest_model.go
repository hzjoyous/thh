//Package guest 模型
package guest

import (
	"thh/app/models"
	"thh/helpers/db"
)

type Guest struct {
	models.BaseModel

	// Put fields in here
	// FIXME()
}

func (guest *Guest) Create() {
	db.SqlDBIns().Create(&guest)
}

func (guest *Guest) Save() (rowsAffected int64) {
	result := db.SqlDBIns().Save(&guest)
	return result.RowsAffected
}

func (guest *Guest) Delete() (rowsAffected int64) {
	result := db.SqlDBIns().Delete(&guest)
	return result.RowsAffected
}
