/*
 * @Author: weihua hu
 * @Date: 2024-12-26 16:17:06
 * @LastEditTime: 2024-12-26 16:17:07
 * @LastEditors: weihua hu
 * @Description:
 */
package initialize

import (
	"sky-take-out-go/global"

	"github.com/go-redis/redis"
)

func InitRedis() {
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "",               // Redis 密码，如果没有设置可为空
		DB:       0,                // 使用默认 DB
	})
}
