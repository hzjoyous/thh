//Package about 模型
package about

import (
	"thh/app/models"
	"thh/helpers/db"
)

type About struct {
	models.BaseModel

	// Put fields in here
	// FIXME()
}

func (about *About) Create() {
	db.SqlDBIns().Create(&about)
}

func (about *About) Save() (rowsAffected int64) {
	result := db.SqlDBIns().Save(&about)
	return result.RowsAffected
}

func (about *About) Delete() (rowsAffected int64) {
	result := db.SqlDBIns().Delete(&about)
	return result.RowsAffected
}
