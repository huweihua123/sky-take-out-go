/*
 * @Author: weihua hu
 * @Date: 2024-12-15 21:10:35
 * @LastEditTime: 2024-12-15 21:11:25
 * @LastEditors: weihua hu
 * @Description:
 */
package initialize

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"sky-take-out-go/global"
)

func InitDB() {
	// 连接数据库
	dsn := "root:123456hwh@tcp(127.0.0.1:3306)/sky-take-out?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)

	// 全局模式
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
}
