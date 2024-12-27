/*
 * @Author: weihua hu
 * @Date: 2024-12-16 00:12:27
 * @LastEditTime: 2024-12-27 01:25:56
 * @LastEditors: weihua hu
 * @Description:
 */
package initialize

import (
	"sky-take-out-go/config"
	"sky-take-out-go/global"
)

func InitConfig() {
	global.ServerConfig.JWTInfo = config.JwtProperties{
		AdminSecretKey: "adminSecretKey",
		AdminTtl:       60 * 60 * 24 * 7,
		AdminTokenName: "token",
		UserSecretKey:  "userSecretKey",
		UserTtl:        60 * 60 * 24 * 7,
		UserTokenName:  "Authorization",
	}
}
