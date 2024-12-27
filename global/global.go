/*
 * @Author: weihua hu
 * @Date: 2024-12-15 21:06:51
 * @LastEditTime: 2024-12-26 21:21:03
 * @LastEditors: weihua hu
 * @Description:
 */
package global

import (
	"sky-take-out-go/config"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	DB               *gorm.DB
	ServerConfig     config.ServerConfig
	WeChatConfig     config.WeChatConfig
	WeChatProperties config.WeChatProperties
	RedisClient      *redis.Client
	JWTConfig        config.JwtProperties
)
