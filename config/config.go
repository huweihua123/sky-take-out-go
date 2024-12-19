/*
 * @Author: weihua hu
 * @Date: 2024-12-16 00:15:52
 * @LastEditTime: 2024-12-16 00:15:53
 * @LastEditors: weihua hu
 * @Description:
 */
package config

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type ServerConfig struct {
	JWTInfo JWTConfig `mapstructure:"jwt" json:"jwt"`
}
