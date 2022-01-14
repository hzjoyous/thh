package v1

import (
	"thh/helpers/response"

	"github.com/gin-gonic/gin"
)

func About(c *gin.Context) {
	response.JSON(c, "删除失败，请稍后尝试~")
}
