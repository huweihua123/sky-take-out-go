/*
 * @Author: weihua hu
 * @Date: 2024-12-15 14:56:09
 * @LastEditTime: 2024-12-15 15:08:59
 * @LastEditors: weihua hu
 * @Description:
 */
package dto

type EmployeeDTO struct {
	ID       int64  `json:"id" gorm:"column:id"`
	Username string `json:"username" gorm:"column:username"`
	Name     string `json:"name" gorm:"column:name"`
	Phone    string `json:"phone" gorm:"column:phone"`
	Sex      string `json:"sex" gorm:"column:sex"`
	IDNumber string `json:"id_number" gorm:"column:id_number"`
}

type EmployeeLoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
