/*
 * @Author: weihua hu
 * @Date: 2024-12-25 13:32:48
 * @LastEditTime: 2024-12-25 13:48:10
 * @LastEditors: weihua hu
 * @Description:
 */
package dto

type SetmealPageQueryDTO struct {
	Name       string `json:"name" form:"name"`
	Page       int    `json:"page" form:"page" binding:"required"`
	PageSize   int    `json:"pageSize" form:"pageSize" binding:"required"`
	Status     *int   `json:"status" form:"status,omitempty"`
	CategoryId int64  `json:"categoryId" form:"categoryId"`
}
