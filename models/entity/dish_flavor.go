/*
 * @Author: weihua hu
 * @Date: 2024-12-21 16:28:28
 * @LastEditTime: 2024-12-21 19:36:07
 * @LastEditors: weihua hu
 * @Description:
 */
package entity

type DishFlavor struct {
	BaseModel
	DishId int64  `gorm:"column:dish_id;type:bigint;not null" json:"dishId"`
	Name   string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Value  string `gorm:"column:value;type:varchar(255);not null" json:"value"`
}
