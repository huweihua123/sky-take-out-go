/*
 * @Author: weihua hu
 * @Date: 2024-12-26 16:40:00
 * @LastEditTime: 2024-12-27 00:57:52
 * @LastEditors: weihua hu
 * @Description:
 */
package router

import (
	api "sky-take-out-go/api/user"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("/login", api.Redirect)
		UserRouter.GET("/callback", api.Login)
		UserRouter.GET("wechat", api.CheckSignature)

	}
}
