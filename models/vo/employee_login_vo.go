/*
 * @Author: weihua hu
 * @Date: 2024-12-16 00:20:58
 * @LastEditTime: 2024-12-16 00:20:59
 * @LastEditors: weihua hu
 * @Description:
 */
package vo

type EmployeeLoginVO struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Token    string `json:"token"`
}
