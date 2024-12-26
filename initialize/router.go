/*
 * @Author: weihua hu
 * @Date: 2024-12-15 15:03:17
 * @LastEditTime: 2024-12-27 01:05:06
 * @LastEditors: weihua hu
 * @Description:
 */
package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"sky-take-out-go/global"
	"sky-take-out-go/middlewares"
	"sky-take-out-go/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	// 静态文件服务（提供 QR 目录访问）
	Router.Static("/static", "./static")
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	//配置跨域
	// Router.Use(middlewares.Cors())
	userMiddleware := middlewares.NewJwtUserMiddleware(&global.ServerConfig.JWTInfo)
	adminMiddleware := middlewares.NewJwtAdminMiddleware(&global.ServerConfig.JWTInfo)

	// Admin 路由组
	adminGroup := Router.Group("/admin")

	adminGroup.Use(adminMiddleware.Handle())
	router.InitEmployeeRouter(adminGroup)
	router.InitCategoryRouter(adminGroup)
	router.InitDishRouter(adminGroup)
	router.InitSetmealRouter(adminGroup)

	// userGroup := Router.Group("/user")
	// userGroup.Use(userMiddleware.Handle())
	// router.InitUserRouter(userGroup)
	// router.InitAddressBookRouter(userGroup)

	// User公开路由组
	userPublicGroup := Router.Group("/user")
	{
		router.InitUserRouter(userPublicGroup) // 处理登录等公开接口
	}

	// User认证路由组
	userAuthGroup := Router.Group("/user")
	userAuthGroup.Use(userMiddleware.Handle())
	{
		router.InitAddressBookRouter(userAuthGroup)
	}

	return Router
}
