/*
 * @Author: weihua hu
 * @Date: 2024-12-26 16:10:17
 * @LastEditTime: 2024-12-26 22:44:18
 * @LastEditors: weihua hu
 * @Description:
 */
package user

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"net/url"
	result "sky-take-out-go/common/result"
	"sky-take-out-go/global"
	"sky-take-out-go/middlewares"
	"sky-take-out-go/models"
	"sky-take-out-go/models/dto"
	"sky-take-out-go/models/vo"
	service "sky-take-out-go/service/user"
	"sky-take-out-go/utils"
	"sort"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

const WX_LOGIN = "https://api.weixin.qq.com/sns/jscode2session"
const TOKEN = "111"

// 配置公众号的token
func CheckSignature(c *gin.Context) {
	// 获取查询参数中的签名、时间戳和随机数
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")
	// 创建包含令牌、时间戳和随机数的字符串切片
	tmpArr := []string{TOKEN, timestamp, nonce}
	// 对切片进行字典排序
	sort.Strings(tmpArr)
	// 将排序后的元素拼接成单个字符串
	tmpStr := ""
	for _, v := range tmpArr {
		tmpStr += v
	}
	// 对字符串进行SHA-1哈希计算
	tmpHash := sha1.New()
	tmpHash.Write([]byte(tmpStr))
	tmpStr = fmt.Sprintf("%x", tmpHash.Sum(nil))
	fmt.Println(tmpStr)
	fmt.Println(signature)
	// 将计算得到的签名与请求中提供的签名进行比较，并根据结果发送相应的响应
	if tmpStr == signature {
		c.String(200, echostr)
		// 修改Redis Set调用
		global.RedisClient.Set("library:token", tmpStr, 7*24*time.Hour)
	} else {
		c.String(403, "签名验证失败 "+timestamp)
	}
}

func Login(ctx *gin.Context) {
	var userLoginDTO dto.UserLoginDTO

	code := ctx.Query("code")
	if code == "" {
		ctx.JSON(http.StatusBadRequest, result.Error("code is required"))
		return
	}
	userLoginDTO.Code = code

	// 调用服务层
	user, err := service.WxLogin(userLoginDTO)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	j := middlewares.NewUserJWT()
	claims := models.CustomClaims{
		ID:       user.ID,
		NickName: user.Name,
		Role:     "user",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: global.JWTConfig.UserTtl,
			Issuer:    "sky-take-out-go",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Error("生成令牌失败"))
		return
	}

	var userLoginVO vo.UserLoginVO = vo.UserLoginVO{
		ID:     user.ID,
		Openid: user.Openid,
		Token:  token,
	}

	ctx.JSON(http.StatusOK, result.Success(userLoginVO))

}

func Redirect(c *gin.Context) {
	state := utils.GenerateRandomState(16) //防止跨站请求伪造攻击 增加安全性
	CallbackDomain := "8nq03916ac13.vicp.fun"
	redirectURL := url.QueryEscape(fmt.Sprintf("https://%s/user/user/callback", CallbackDomain))
	fmt.Println(redirectURL)
	wechatLoginURL := fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&state=%s&scope=snsapi_userinfo#wechat_redirect", "wxd1b10d6051741ed5", redirectURL, state)
	wechatLoginURL, _ = url.QueryUnescape(wechatLoginURL)
	// 生成二维码
	qrCode, err := qrcode.Encode(wechatLoginURL, qrcode.Medium, 256)
	if err != nil {
		// 错误处理
		c.String(http.StatusInternalServerError, "Error generating QR code")
		return
	}
	// 将二维码图片作为响应返回给用户
	c.Header("Content-Type", "image/png")
	c.Writer.Write(qrCode)
}
