/*
 * @Author: weihua hu
 * @Date: 2024-12-26 20:24:51
 * @LastEditTime: 2024-12-28 00:02:52
 * @LastEditors: weihua hu
 * @Description:
 */
package router

import (
	api "sky-take-out-go/api/user/address_book"

	"github.com/gin-gonic/gin"
)

func InitAddressBookRouter(Router *gin.RouterGroup) {
	AddressBookRouter := Router.Group("addressBook")
	{
		AddressBookRouter.POST("", api.Create)
		AddressBookRouter.PUT("default", api.SetDefault)
		AddressBookRouter.GET("default", api.GetDefault)
		AddressBookRouter.PUT("", api.Update)
		AddressBookRouter.DELETE("", api.Delete)
		AddressBookRouter.GET(":id", api.GetById)
		AddressBookRouter.GET("list", api.List)
	}
}
