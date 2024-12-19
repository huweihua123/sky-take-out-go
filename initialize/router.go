/*
 * @Author: weihua hu
 * @Date: 2024-12-15 15:03:17
 * @LastEditTime: 2024-12-20 00:30:00
 * @LastEditors: weihua hu
 * @Description:
 */
package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"sky-take-out-go/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	//配置跨域
	// Router.Use(middlewares.Cors())

	// Admin 路由组
	adminGroup := Router.Group("/admin")
	router.InitEmployeeRouter(adminGroup)
	router.InitCategoryRouter(adminGroup)

	// userGroup := Router.Group("/user")

	return Router
}
