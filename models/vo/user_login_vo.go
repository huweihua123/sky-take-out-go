/*
 * @Author: weihua hu
 * @Date: 2024-12-26 16:28:24
 * @LastEditTime: 2024-12-26 16:28:25
 * @LastEditors: weihua hu
 * @Description:
 */
package vo

type UserLoginVO struct {
	ID     int64  `json:"id"`
	Openid string `json:"openid"`
	Token  string `json:"token"`
}
