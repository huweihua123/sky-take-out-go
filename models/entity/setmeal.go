/*
 * @Author: weihua hu
 * @Date: 2024-12-22 21:08:58
 * @LastEditTime: 2024-12-22 22:52:55
 * @LastEditors: weihua hu
 * @Description:
 */
package entity

type Setmeal struct {
	BaseModel
	CategoryId  int64   `json:"categoryId" gorm:"not null"`
	Name        string  `json:"name" gorm:"size:255;not null"`
	Price       float64 `json:"price" gorm:"not null"`
	Status      int     `json:"status" gorm:"not null"`
	Description string  `json:"description" gorm:"size:255"`
	CreateUser  int64   `json:"create_user" gorm:"not null"`
	UpdateUser  int64   `json:"update_user" gorm:"not null"`
	Image       string  `json:"image" gorm:"size:255"`
}
