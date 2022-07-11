package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ERROR   = 0
	SUCCESS = 1
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"dataRep"`
	Msg  string `json:"msg"`
}

func Result(code int, data any, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func FailWithDetailed(data any, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
