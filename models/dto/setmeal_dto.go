/*
 * @Author: weihua hu
 * @Date: 2024-12-22 21:13:12
 * @LastEditTime: 2024-12-23 16:09:35
 * @LastEditors: weihua hu
 * @Description:
 */
package dto

import "sky-take-out-go/models/entity"

type SetMealDTO struct {
	CategoryId    int64                `json:"categoryId" form:"categoryId" binding:"required"`
	Description   string               `json:"description" form:"description"`
	Id            int64                `json:"id" form:"id"`
	Image         string               `json:"image" form:"image" binging:"required"`
	Name          string               `json:"name" form:"name" binding:"required"`
	Price         float64              `json:"price" form:"price" binding:"required"`
	Status        int                  `json:"status" form:"status" ` // 0: 停售, 1: 起售
	SetmealDishes []entity.SetmealDish `json:"setMealDishes" form:"setMealDishes" binding:"required"`
}
