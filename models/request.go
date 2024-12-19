/*
 * @Author: weihua hu
 * @Date: 2024-12-16 00:10:00
 * @LastEditTime: 2024-12-16 00:37:42
 * @LastEditors: weihua hu
 * @Description:
 */
package models

import (
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	ID       int64
	NickName string
	jwt.StandardClaims
}
