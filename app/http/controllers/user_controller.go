package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"runtime"
	"thh/app/models/Users"
	"thh/arms/jwt"
	"time"
)

const (
	expireTime = time.Second * 86400 * 7
)

// Register
// @todo user表增加验证字段
// 创建后验证码存入redis，发ggtgtrtrftrftr送邮件。
// 邮件 附有 url?code=xxx
// 验证后更新验证字段
// 清除验证码
func Register(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	type request struct {
		Email    string `json:"email" binging:"required"`
		Username string `json:"userName"  binding:"required"`
		Password string `json:"passWord"  binding:"required"`
		NickName string `json:"nickName" gorm:"default:'QMPlusUser'"`
	}
	var r request
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userEntity := Users.MakeUser(r.Username, r.Password, r.Email)
	err := userEntity.Create()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "创建失败",
		})
		return
	}

	token, err := jwt.CreateNewToken(userEntity.ID, expireTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "创建失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"token":   token,
	})
}

func Login(c *gin.Context) {
	type request struct {
		Username string `json:"userName"  binding:"required"`
		Password string `json:"passWord"  binding:"required"`
	}
	var r request
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userEntity, err := Users.Verify(r.Username, r.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "登录失败",
		})
		return
	}

	token, err := jwt.CreateNewToken(userEntity.ID, expireTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "创建失败",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"token":   token,
	})

}

func GetUserInfo(request Request) Response {
	type RequestData struct {
		UserId uint64 `json:"userId"  binding:"required"`
	}
	var requestData RequestData
	err := request.Context.BindQuery(&requestData)
	if err != nil {
		return FailResponse(err.Error())
	}

	userEntity, err := Users.GetById(requestData.UserId)
	if err != nil {
		return FailResponse(err.Error())
	}
	return SuccessResponse(userEntity)
}

func UserInfoV2(request Request) Response {
	userEntity, err := request.GetUser()
	if err != nil {
		return FailResponse("账号异常" + err.Error())
	}
	return SuccessResponse(userEntity)
}

func GetUseMem(Request) Response {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return SuccessResponse(cast.ToString(m.Alloc/1024/8) + "kb")
}
