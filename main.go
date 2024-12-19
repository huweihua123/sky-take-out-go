/*
 * @Author: weihua hu
 * @Date: 2024-12-15 14:52:05
 * @LastEditTime: 2024-12-15 23:02:05
 * @LastEditors: weihua hu
 * @Description:
 */
package main

import (
	"fmt"
	"sky-take-out-go/initialize"

	"go.uber.org/zap"
)

func main() {
	initialize.InitLogger()
	initialize.InitDB()
	Router := initialize.Routers()

	Port := 8084
	zap.S().Debugf("启动服务器, 端口： %d", Port)

	// 启动gin服务
	if err := Router.Run(fmt.Sprintf(":%d", Port)); err != nil {
		zap.S().Panic("启动失败:", err.Error())
	}
}
