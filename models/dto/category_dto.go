/*
 * @Author: weihua hu
 * @Date: 2024-12-20 00:31:07
 * @LastEditTime: 2024-12-20 14:59:23
 * @LastEditors: weihua hu
 * @Description:
 */
package dto

type CategoryDTO struct {
	ID   int64  `json:"id,string"`
	Name string `json:"name" binding:"required"`
	Sort int    `json:"sort" binding:"required"`
	Type int    `json:"type" binding:"required"`
}
