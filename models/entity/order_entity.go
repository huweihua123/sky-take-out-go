/*
 * @Author: weihua hu
 * @Date: 2024-12-25 14:05:59
 * @LastEditTime: 2024-12-25 14:06:00
 * @LastEditors: weihua hu
 * @Description:
 */
package entity

type OrderDetail struct {
	BaseModel
	Name       string  `json:"name" gorm:"column:name;type:varchar(100);not null"`      // 名称
	OrderId    int64   `json:"orderId" gorm:"column:order_id;type:bigint;not null"`     // 订单id
	DishId     int64   `json:"dishId" gorm:"column:dish_id;type:bigint"`                // 菜品id
	SetmealId  int64   `json:"setmealId" gorm:"column:setmeal_id;type:bigint"`          // 套餐id
	DishFlavor string  `json:"dishFlavor" gorm:"column:dish_flavor;type:varchar(100)"`  // 口味
	Number     int     `json:"number" gorm:"column:number;type:int;not null"`           // 数量
	Amount     float64 `json:"amount" gorm:"column:amount;type:decimal(10,2);not null"` // 金额
	Image      string  `json:"image" gorm:"column:image;type:varchar(255)"`             // 图片
}
