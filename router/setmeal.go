/*
 * @Author: weihua hu
 * @Date: 2024-12-22 21:12:14
 * @LastEditTime: 2024-12-25 13:30:55
 * @LastEditors: weihua hu
 * @Description:
 */
package router

import (
	api "sky-take-out-go/api/setmeal"

	"github.com/gin-gonic/gin"
)

func InitSetmealRouter(Router *gin.RouterGroup) {
	SetmealRouter := Router.Group("setmeal")
	{
		SetmealRouter.POST("", api.Create)
		SetmealRouter.PUT("", api.Update)
		SetmealRouter.DELETE("", api.Delete)
		SetmealRouter.GET(":id", api.GetById)
		SetmealRouter.GET("page", api.Page)
		// SetmealRouter.GET("list", api.List)
		SetmealRouter.POST("status/:status", api.StartOrStop)
	}
}
