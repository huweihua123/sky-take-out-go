/*
 * @Author: weihua hu
 * @Date: 2024-12-15 23:08:15
 * @LastEditTime: 2024-12-15 23:09:17
 * @LastEditors: weihua hu
 * @Description:
 */
package main

import (
	"fmt"
	"sky-take-out-go/utils"
)

func main() {
	password := "testpassword"
	encryptedPassword := utils.MD5Encrypt(password)
	fmt.Println(encryptedPassword)
}
