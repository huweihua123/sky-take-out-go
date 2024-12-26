/*
 * @Author: weihua hu
 * @Date: 2024-12-15 15:05:03
 * @LastEditTime: 2024-12-20 18:30:41
 * @LastEditors: weihua hu
 * @Description:
 */
package router

import (
	api "sky-take-out-go/api/employee"

	"github.com/gin-gonic/gin"
)

func InitEmployeeRouter(Router *gin.RouterGroup) {
	EmployeeRouter := Router.Group("employee")
	{
		EmployeeRouter.POST("login", api.Login)                 // 员工登录
		EmployeeRouter.POST("save", api.Save)                   // 新增员工
		EmployeeRouter.POST("/status/:status", api.StartOrStop) // 启用或禁用员工账号
		EmployeeRouter.PUT("update", api.Update)                // 启用或禁用员工账号
		EmployeeRouter.GET(":id", api.GetById)                  // 获取员工信息
		EmployeeRouter.GET("page", api.Page)                    // 分页员工信息
	}
}
