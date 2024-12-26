/*
 * @Author: weihua hu
 * @Date: 2024-12-26 17:05:30
 * @LastEditTime: 2024-12-26 17:05:31
 * @LastEditors: weihua hu
 * @Description:
 */
package utils

import "math/rand"

func GenerateRandomState(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
