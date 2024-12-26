/*
 * @Author: weihua hu
 * @Date: 2024-12-20 12:56:31
 * @LastEditTime: 2024-12-20 12:56:32
 * @LastEditors: weihua hu
 * @Description:
 */
package dto

type CategoryPageQueryDTO struct {
	Name     string `json:"name" form:"name"`
	Page     int    `json:"page" form:"page" binding:"required"`
	PageSize int    `json:"pageSize" form:"pageSize" binding:"required"`
	Type     int    `json:"type" form:"type"`
}
