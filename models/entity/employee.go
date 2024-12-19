/*
 * @Author: weihua hu
 * @Date: 2024-12-15 14:41:10
 * @LastEditTime: 2024-12-16 14:59:25
 * @LastEditors: weihua hu
 * @Description:
 */
package entity

type Employee struct {
	BaseModel
	Username   string `json:"username" gorm:"size:255;not null"`
	Name       string `json:"name" gorm:"size:255;not null"`
	Password   string `json:"password" gorm:"size:255;not null"`
	Phone      string `json:"phone" gorm:"size:20"`
	Sex        string `json:"sex" gorm:"size:10"`
	IDNumber   string `json:"id_number" gorm:"size:20"`
	Status     int    `json:"status" gorm:"not null"`
	CreateUser int64  `json:"create_user"`
	UpdateUser int64  `json:"update_user"`
}

func (Employee) TableName() string {
	return "employee"
}
