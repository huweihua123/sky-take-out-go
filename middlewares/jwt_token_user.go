/*
 * @Author: weihua hu
 * @Date: 2024-12-26 22:08:55
 * @LastEditTime: 2024-12-27 01:32:00
 * @LastEditors: weihua hu
 * @Description:
 */
package middlewares

import (
	"net/http"
	"sky-take-out-go/config"

	"github.com/gin-gonic/gin"
)

const (
	CtxUserIDKey = "currentUserId"
)

type JwtUserMiddleware struct {
	JwtProps *config.JwtProperties
}

func NewJwtUserMiddleware(props *config.JwtProperties) *JwtUserMiddleware {
	return &JwtUserMiddleware{
		JwtProps: props,
	}
}

func (j *JwtUserMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(j.JwtProps.UserTokenName)

		if token == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"msg": "请登录",
			})
			c.Abort()
			return
		}

		j := NewUserJWT()

		claims, err := j.ParseToken(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 解析用户ID
		c.Set(CtxUserIDKey, claims.ID)
		c.Next()
	}
}

func GetCurrentUserID(c *gin.Context) int64 {
	value, exists := c.Get(CtxUserIDKey)
	if !exists {
		return 0
	}
	return value.(int64)
}
