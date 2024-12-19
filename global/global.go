/*
 * @Author: weihua hu
 * @Date: 2024-12-15 21:06:51
 * @LastEditTime: 2024-12-16 00:15:18
 * @LastEditors: weihua hu
 * @Description:
 */
package global

import (
	"sky-take-out-go/config"

	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
)
