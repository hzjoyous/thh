package guest

import (
	"thh/helpers/app"
	"thh/helpers/db"
	"thh/helpers/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (guest Guest) {
	db.SqlDBIns().Where("id", idstr).First(&guest)
	return
}

func GetBy(field, value string) (guest Guest) {
	db.SqlDBIns().Where("? = ?", field, value).First(&guest)
	return
}

func All() (guests []Guest) {
	db.SqlDBIns().Find(&guests)
	return 
}

func IsExist(field, value string) bool {
	var count int64
	db.SqlDBIns().Model(Guest{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (guests []Guest, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		db.SqlDBIns().Model(Guest{}),
		&guests,
		app.V1URL(db.TableName(&Guest{})),
		perPage,
	)
	return
}