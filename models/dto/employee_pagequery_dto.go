/*
 * @Author: weihua hu
 * @Date: 2024-12-19 22:38:17
 * @LastEditTime: 2024-12-19 23:19:03
 * @LastEditors: weihua hu
 * @Description:
 */
package dto

type EmployeePageQueryDTO struct {
	Name     string `json:"name" form:"name"`
	Page     int    `json:"page" form:"page" binding:"required"`
	PageSize int    `json:"pageSize" form:"pageSize" binding:"required"`
}
