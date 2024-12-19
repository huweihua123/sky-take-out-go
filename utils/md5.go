/*
 * @Author: weihua hu
 * @Date: 2024-12-15 21:24:51
 * @LastEditTime: 2024-12-15 21:24:52
 * @LastEditors: weihua hu
 * @Description:
 */
package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5Encrypt 对字符串进行MD5加密
func MD5Encrypt(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
