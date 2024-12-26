/*
 * @Author: weihua hu
 * @Date: 2024-12-20 00:09:45
 * @LastEditTime: 2024-12-20 15:13:58
 * @LastEditors: weihua hu
 * @Description:
 */
package entity

type Category struct {
	BaseModel
	Type       int    `json:"type" gorm:"not null"`
	Name       string `json:"name" gorm:"size:255;not null"`
	Sort       int    `json:"sort" gorm:"not null"`
	Status     int    `json:"status" gorm:"not null"`
	CreateUser int64  `json:"create_user" gorm:"not null"`
	UpdateUser int64  `json:"update_user" gorm:"not null"`
}
