/*
 * @Author: weihua hu
 * @Date: 2024-12-25 20:41:04
 * @LastEditTime: 2024-12-26 16:09:40
 * @LastEditors: weihua hu
 * @Description:
 */
package initialize

import "sky-take-out-go/global"

func InitWx() {
	global.WeChatConfig.AppID = "wxd1b10d6051741ed5"
	global.WeChatConfig.AppSecret = "de3e20ad97b00959db7acb9dbccecdca"
	global.WeChatConfig.RedirectURL = "https://8nq03916ac13.vicp.fun/user/wechat/callback"
}
