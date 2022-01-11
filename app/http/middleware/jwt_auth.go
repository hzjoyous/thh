package middleware

import (
	"github.com/gin-gonic/gin"
	"thh/app/service/response"
	"thh/helpers"
	"thh/helpers/jwt"
	"time"
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
		claims, err := jwt.UseJWT().ParseToken(token)
		// 验证失败
		if err != nil {
			errorMsg := err.Error()
			if err == jwt.TokenExpired {
				errorMsg = "授权已过期"
			}
			response.FailWithDetailed(gin.H{"reload": true}, errorMsg, c)
			c.Abort()
			return
		}
		// 如果过期时间减去当前时间小于缓冲时间
		// 说明令牌即将过期
		// 需要更新为新的令牌
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + 86400
			newToken, _ := jwt.UseJWT().CreateToken(*claims)
			newClaims, _ := jwt.UseJWT().ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", helpers.ToString(newClaims.ExpiresAt))
		}
		c.Set("userId", claims.UserId)
		c.Next()
	}
}
