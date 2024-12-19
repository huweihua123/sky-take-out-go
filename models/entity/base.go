/*
 * @Author: weihua hu
 * @Date: 2024-12-15 15:00:21
 * @LastEditTime: 2024-12-16 15:13:54
 * @LastEditors: weihua hu
 * @Description:
 */
package entity

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int64          `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"column:create_time"`
	UpdatedAt time.Time      `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	IsDeleted bool
}
