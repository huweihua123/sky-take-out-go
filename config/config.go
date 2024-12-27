/*
 * @Author: weihua hu
 * @Date: 2024-12-16 00:15:52
 * @LastEditTime: 2024-12-26 21:21:57
 * @LastEditors: weihua hu
 * @Description:
 */
package config

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type ServerConfig struct {
	JWTInfo JwtProperties `mapstructure:"jwt" json:"jwt"`
}

type WeChatConfig struct {
	AppID       string
	AppSecret   string
	RedirectURL string
}

type WeChatProperties struct {
	AppID                 string `json:"appid"`
	Secret                string `json:"secret"`
	MchID                 string `json:"mchid"`
	MchSerialNo           string `json:"mchSerialNo"`
	PrivateKeyFilePath    string `json:"privateKeyFilePath"`
	ApiV3Key              string `json:"apiV3Key"`
	WeChatPayCertFilePath string `json:"weChatPayCertFilePath"`
	NotifyUrl             string `json:"notifyUrl"`
	RefundNotifyUrl       string `json:"refundNotifyUrl"`
}

type JwtProperties struct {
	// 管理端员工生成jwt令牌相关配置
	AdminSecretKey string `json:"adminSecretKey"`
	AdminTtl       int64  `json:"adminTtl"`
	AdminTokenName string `json:"adminTokenName"`

	// 用户端微信用户生成jwt令牌相关配置
	UserSecretKey string `json:"userSecretKey"`
	UserTtl       int64  `json:"userTtl"`
	UserTokenName string `json:"userTokenName"`
}
