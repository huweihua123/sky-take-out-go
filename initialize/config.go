/*
 * @Author: weihua hu
 * @Date: 2024-12-16 00:12:27
 * @LastEditTime: 2024-12-16 00:16:29
 * @LastEditors: weihua hu
 * @Description:
 */
package initialize

import (
	"sky-take-out-go/config"
	"sky-take-out-go/global"
)

func InitConfig() {
	global.ServerConfig.JWTInfo = config.JWTConfig{
		SigningKey: "sky-take-out",
	}
}
