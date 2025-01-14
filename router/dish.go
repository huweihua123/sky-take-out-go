/*
 * @Author: weihua hu
 * @Date: 2024-12-21 16:53:13
 * @LastEditTime: 2024-12-27 23:59:05
 * @LastEditors: weihua hu
 * @Description:
 */
package router

import (
	api "sky-take-out-go/api/admin/dish"

	"github.com/gin-gonic/gin"
)

func InitDishRouter(Router *gin.RouterGroup) {
	DishRouter := Router.Group("dish")
	{
		DishRouter.POST("", api.Create)
		DishRouter.PUT("", api.Update)
		DishRouter.DELETE("", api.Delete)
		DishRouter.GET(":id", api.GetById)
		// DishRouter.GET("list", api.List)
		DishRouter.POST("status/:status", api.StartOrStop)

	}
}
