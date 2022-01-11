package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/utils"
	"net/http"
	"runtime"
	"thh/app/models/user"
	"thh/helpers/jwt"
)

// Register
// @todo user表增加验证字段
// 创建后验证码存入redis，发ggtgtrtrftrftr送邮件。
// 邮件 附有 url?code=xxx
// 验证后更新验证字段
// 清除验证码
func Register(c *gin.Context) {
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

	userEntity := user.MakeUser(r.Username, r.Password, r.Email)
	err := userEntity.Create()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "创建失败",
		})
		return
	}

	token, err := jwt.CreateNewToken(userEntity.ID, userEntity.Username)
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

	userEntity, err := user.UserRepository().Verify(r.Username, r.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "登录失败",
		})
		return
	}

	token, err := jwt.CreateNewToken(userEntity.ID, userEntity.Username)
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

	user, err := user.UserRepository().GetById(requestData.UserId)
	if err != nil {
		return FailResponse(err.Error())
	}
	return SuccessResponse(user)
}

func UserInfoV2(request Request) Response {
	user, err := request.GetUser()
	if err != nil {
		return FailResponse("账号异常" + err.Error())
	}
	return SuccessResponse(user)
}

func GetUseMem(Request) Response {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return SuccessResponse(utils.ToString(m.Alloc/1024/8) + "kb")
}
