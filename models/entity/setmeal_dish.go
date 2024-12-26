/*
 * @Author: weihua hu
 * @Date: 2024-12-22 21:30:54
 * @LastEditTime: 2024-12-22 21:49:03
 * @LastEditors: weihua hu
 * @Description:
 */
package entity

type SetmealDish struct {
	BaseModel
	SetmealId int64   `json:"setmealId"`
	DishId    int64   `json:"dishId"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Copies    int     `json:"copies"`
}
