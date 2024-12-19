/*
 * @Author: weihua hu
 * @Date: 2024-12-19 22:43:32
 * @LastEditTime: 2024-12-19 23:08:19
 * @LastEditors: weihua hu
 * @Description:
 */
package result

type PageResult struct {
	Total   int         `json:"total"`
	Records interface{} `json:"records"`
}
