/*
 * @Author: weihua hu
 * @Date: 2024-12-26 16:40:00
 * @LastEditTime: 2024-12-28 00:00:11
 * @LastEditors: weihua hu
 * @Description:
 */
package router

import (
	api "sky-take-out-go/api/user/user"

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
