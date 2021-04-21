package jwt

import (
	"gin-app/pkg/e"
	"gin-app/pkg/util"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS

		authHeader := c.Request.Header.Get("Authorization") // 获取请求头中的数据
		parts := strings.SplitN(authHeader, " ", 2)

		if !(len(parts) == 2 && parts[0] == "Bearer") {
			code = e.INVALID_PARAMS
		} else {
			// parts[1] 是获取到的 token
			claims, err := util.ParseToken(parts[1])
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}

			// 将当前请求的 username 信息保存到请求的上下文c上
			c.Set("userName", claims.Username)
			c.Set("userId", claims.Id)
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
