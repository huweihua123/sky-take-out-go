/*
 * @Author: weihua hu
 * @Date: 2024-12-15 23:31:23
 * @LastEditTime: 2024-12-26 20:25:26
 * @LastEditors: weihua hu
 * @Description:
 */
package main

import (
	"log"
	"os"
	"sky-take-out-go/models/entity"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func main() {
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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	_ = db.AutoMigrate(&entity.AddressBook{})
	// _ = db.AutoMigrate(&entity.Employee{}, &entity.Category{}, &entity.Dish{}, &entity.DishFlavor{})

}
