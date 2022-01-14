package about

import (
	"thh/helpers/app"
	"thh/helpers/db"
	"thh/helpers/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (about About) {
	db.SqlDBIns().Where("id", idstr).First(&about)
	return
}

func GetBy(field, value string) (about About) {
	db.SqlDBIns().Where("? = ?", field, value).First(&about)
	return
}

func All() (abouts []About) {
	db.SqlDBIns().Find(&abouts)
	return 
}

func IsExist(field, value string) bool {
	var count int64
	db.SqlDBIns().Model(About{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (abouts []About, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		db.SqlDBIns().Model(About{}),
		&abouts,
		app.V1URL(db.TableName(&About{})),
		perPage,
	)
	return
}