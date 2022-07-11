package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thh/app/models/Users"
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
	userInfo Users.User
	Context  *gin.Context
}

func (r *Request) GetUser() (Users.User, error) {
	if r.userSet != false {
		return r.userInfo, nil
	}
	user, err := Users.GetById(r.UserId)
	if err != nil {
		return r.userInfo, err
	}
	r.userSet = true
	r.userInfo = user
	return r.userInfo, nil
}

type Response struct {
	Code int
	Data any
}

func BuildResponse(code int, data any) Response {
	return Response{code, data}
}

func SuccessResponse(data any) Response {
	return BuildResponse(http.StatusOK,
		map[string]any{
			"msg":     nil,
			"dataRep": data,
			"code":    SUCCESS,
		},
	)
}

func FailResponse(msg string) Response {
	return BuildResponse(http.StatusOK,
		map[string]any{
			"msg":     msg,
			"dataRep": nil,
			"code":    FAIL,
		},
	)
}
