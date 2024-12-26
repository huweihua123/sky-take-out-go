/*
 * @Author: weihua hu
 * @Date: 2024-12-25 13:18:08
 * @LastEditTime: 2024-12-25 13:21:21
 * @LastEditors: weihua hu
 * @Description:
 */
package vo

import "sky-take-out-go/models/entity"

type SetMealVO struct {
	Id            int64
	CategoryId    int64
	CategoryName  string
	Description   string
	Image         string
	Name          string
	Price         float64
	Status        int
	SetMealDishes []entity.SetmealDish
}
