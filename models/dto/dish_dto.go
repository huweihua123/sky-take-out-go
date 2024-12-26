/*
 * @Author: weihua hu
 * @Date: 2024-12-21 16:34:46
 * @LastEditTime: 2024-12-21 16:57:38
 * @LastEditors: weihua hu
 * @Description:
 */
package dto

import "sky-take-out-go/models/entity"

// DishDTO 菜品数据传输对象
type DishDTO struct {
	ID          int64               `json:"id,omitempty"`                  // 菜品id
	Name        string              `json:"name" binding:"required"`       // 菜品名称
	CategoryId  int64               `json:"categoryId" binding:"required"` // 分类id
	Price       float64             `json:"price" binding:"required"`      // 菜品价格
	Image       string              `json:"image" binding:"required"`      // 菜品图片路径
	Description string              `json:"description"`                   // 菜品描述
	Status      int                 `json:"status"`                        // 菜品状态：1为起售，0为停售
	Flavors     []entity.DishFlavor `json:"flavors"`                       // 口味
}
