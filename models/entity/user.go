/*
 * @Author: weihua hu
 * @Date: 2024-12-26 14:21:21
 * @LastEditTime: 2024-12-26 14:26:23
 * @LastEditors: weihua hu
 * @Description:
 */
package entity

type User struct {
	BaseModel
	Openid   string `json:"openid" gorm:"column:openid;size:255;not null"`
	Name     string `json:"name" gorm:"size:255;not null"`
	Phone    string `json:"phone" gorm:"size:255;not null"`
	Avatar   string `json:"avatar" gorm:"size:255"`
	IdNumber string `json:"idNumber" gorm:"size:255"`
	Sex      string `json:"sex" gorm:"size:10"`
}
