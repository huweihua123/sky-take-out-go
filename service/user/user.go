/*
 * @Author: weihua hu
 * @Date: 2024-12-26 16:14:03
 * @LastEditTime: 2024-12-26 20:23:20
 * @LastEditors: weihua hu
 * @Description:
 */
package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sky-take-out-go/global"
	"sky-take-out-go/models/dto"
	"sky-take-out-go/models/entity"

	"gorm.io/gorm"
)

func WxLogin(userLoginDTO dto.UserLoginDTO) (entity.User, error) {
	// 获取token
	tokenResp, err := http.Get(fmt.Sprintf(
		"https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		"wxd1b10d6051741ed5",
		"de3e20ad97b00959db7acb9dbccecdca",
		userLoginDTO.Code,
	))
	fmt.Println("error1")

	if err != nil {
		return entity.User{}, fmt.Errorf("请求失败: %v", err)
	}
	defer tokenResp.Body.Close()

	var tokenData struct {
		AccessToken string `json:"access_token"`
		OpenID      string `json:"openid"`
	}
	if err := json.NewDecoder(tokenResp.Body).Decode(&tokenData); err != nil {
		return entity.User{}, fmt.Errorf("解析token失败: %v", err)
	}

	// 获取用户信息

	userResp, err := http.Get(fmt.Sprintf(
		"https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s",
		tokenData.AccessToken,
		tokenData.OpenID,
	))
	if err != nil {
		return entity.User{}, fmt.Errorf("获取用户信息失败: %v", err)
	}

	defer userResp.Body.Close()

	var userData struct {
		OpenID   string `json:"openid"`
		Nickname string `json:"nickname"`
		HeadImg  string `json:"headimgurl"`
	}
	if err := json.NewDecoder(userResp.Body).Decode(&userData); err != nil {
		return entity.User{}, fmt.Errorf("解析用户信息失败: %v", err)
	}

	// 根据 openid 判断用户是否存在
	var existUser entity.User
	if err := global.DB.Where("openid = ?", userData.OpenID).First(&existUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 不存在则创建
			newUser := entity.User{
				Openid: userData.OpenID,
				Name:   userData.Nickname,
				Avatar: userData.HeadImg,
			}
			if err := global.DB.Create(&newUser).Error; err != nil {
				return entity.User{}, fmt.Errorf("创建用户失败: %v", err)
			}
			return newUser, nil
		}
		return entity.User{}, fmt.Errorf("数据库错误: %v", err)
	}

	// 已存在则直接返回
	return existUser, nil
}
