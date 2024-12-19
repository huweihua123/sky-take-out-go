/*
 * @Author: weihua hu
 * @Date: 2024-12-15 22:59:36
 * @LastEditTime: 2024-12-15 22:59:37
 * @LastEditors: weihua hu
 * @Description:
 */
package initialize

import "go.uber.org/zap"

func InitLogger() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
