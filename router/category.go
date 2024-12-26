/*
 * @Author: weihua hu
 * @Date: 2024-12-20 00:25:28
 * @LastEditTime: 2024-12-20 18:30:47
 * @LastEditors: weihua hu
 * @Description:
 */
package router

import (
	api "sky-take-out-go/api/category"

	"github.com/gin-gonic/gin"
)

func InitCategoryRouter(Router *gin.RouterGroup) {
	CategoryRouter := Router.Group("category")
	{
		CategoryRouter.POST("", api.Create)
		CategoryRouter.PUT("", api.Update)
		CategoryRouter.DELETE("", api.Delete)
		CategoryRouter.GET("page", api.Page)
		CategoryRouter.POST("/status/:status", api.StartOrStop)
		CategoryRouter.GET("list", api.List)
	}
}
