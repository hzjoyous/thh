package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thh/app/models/user"
)

const (
	SUCCESS = 1
	FAIL    = 0
)

type BaseController struct {
}
type Request struct {
	TraceId  string
	UserId   uint64
	userSet  bool
	userInfo user.User
	Context  *gin.Context
}

func (r *Request) GetUser() (user.User, error) {
	if r.userSet != false {
		return r.userInfo, nil
	}
	user, err := user.UserRepository().GetById(r.UserId)
	if err != nil {
		return r.userInfo, err
	}
	r.userSet = true
	r.userInfo = user
	return r.userInfo, nil
}

type Response struct {
	Code int
	Data interface{}
}

func BuildResponse(code int, data interface{}) Response {
	return Response{code, data}
}

func SuccessResponse(data interface{}) Response {
	return BuildResponse(http.StatusOK,
		map[string]interface{}{
			"msg":  nil,
			"dataRep": data,
			"code": SUCCESS,
		},
	)
}

func FailResponse(msg string) Response {
	return BuildResponse(http.StatusOK,
		map[string]interface{}{
			"msg":  msg,
			"dataRep": nil,
			"code": FAIL,
		},
	)
}
