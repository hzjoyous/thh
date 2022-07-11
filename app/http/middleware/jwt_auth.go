package middleware

import (
	"github.com/gin-gonic/gin"
	"thh/app/service/response"
	"thh/arms/jwt"
)

// JWTAuth
// 如果未获取到 x-token 则非法登陆
// 如果已经过期 则推出
// 如果只是则在header返回新的token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		userId, newToken, err := jwt.VerifyTokenWithFresh(token)
		if err != nil {
			errorMsg := err.Error()
			if err == jwt.TokenExpired {
				errorMsg = "授权已过期"
			}
			response.FailWithDetailed(gin.H{"reload": true}, errorMsg, c)
			c.Abort()
			return
		}
		if token != newToken {
			c.Header("new-token", newToken)
		}
		c.Set("userId", userId)
		c.Next()
	}
}
