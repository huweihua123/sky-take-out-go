/*
 * @Author: weihua hu
 * @Date: 2024-12-20 00:19:32
 * @LastEditTime: 2024-12-21 16:43:29
 * @LastEditors: weihua hu
 * @Description:
 */
package entity

type Dish struct {
	BaseModel
	Name        string  `json:"name"`
	CategoryId  int64   `json:"category_id"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	Description string  `json:"description"`
	Status      int     `json:"status" gorm:"default:1"` // 1下架 0上架
	CreateUser  int64   `json:"create_user"`
	UpdateUser  int64   `json:"update_user"`
}
