/*
 * @Author: weihua hu
 * @Date: 2024-12-26 22:42:03
 * @LastEditTime: 2024-12-26 22:48:14
 * @LastEditors: weihua hu
 * @Description:
 */
package middlewares

import (
	"fmt"
	"net/http"
	"sky-take-out-go/config"

	"github.com/gin-gonic/gin"
)

type JwtAdminMiddleware struct {
	JwtProps *config.JwtProperties
}

func NewJwtAdminMiddleware(props *config.JwtProperties) *JwtUserMiddleware {
	return &JwtUserMiddleware{
		JwtProps: props,
	}
}

func (j *JwtAdminMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 从请求头获取token
		token := c.GetHeader(j.JwtProps.AdminTokenName)

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
		userId := claims.ID
		fmt.Println(userId)
		// c.Set("currentUserId", userId)
		c.Next()
	}
}
