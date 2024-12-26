/*
 * @Author: weihua hu
 * @Date: 2024-12-26 22:37:15
 * @LastEditTime: 2024-12-26 22:37:16
 * @LastEditors: weihua hu
 * @Description:
 */
package middlewares

import "sky-take-out-go/global"

type JWT struct {
	SigningKey []byte
}

func NewUserJWT() *JWT {
	return &JWT{
		SigningKey: []byte(global.ServerConfig.JWTInfo.UserSecretKey),
	}
}

func NewAdminJWT() *JWT {
	return &JWT{
		SigningKey: []byte(global.ServerConfig.JWTInfo.AdminSecretKey),
	}
}
