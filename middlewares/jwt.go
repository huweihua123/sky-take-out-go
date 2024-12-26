/*
 * @Author: weihua hu
 * @Date: 2024-12-16 00:08:34
 * @LastEditTime: 2024-12-26 22:43:40
 * @LastEditors: weihua hu
 * @Description:
 */
package middlewares

import (
	"errors"
	"sky-take-out-go/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// func JWTAuth() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localSstorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
// 		token := c.Request.Header.Get("x-token")
// 		if token == "" {
// 			c.JSON(http.StatusUnauthorized, map[string]string{
// 				"msg": "请登录",
// 			})
// 			c.Abort()
// 			return
// 		}
// 		j := NewUserJWT()
// 		// parseToken 解析token包含的信息
// 		claims, err := j.ParseToken(token)
// 		if err != nil {
// 			if err == ErrTokenExpired {
// 				if err == ErrTokenExpired {
// 					c.JSON(http.StatusUnauthorized, map[string]string{
// 						"msg": "授权已过期",
// 					})
// 					c.Abort()
// 					return
// 				}
// 			}

// 			c.JSON(http.StatusUnauthorized, "未登陆")
// 			c.Abort()
// 			return
// 		}
// 		c.Set("claims", claims)
// 		c.Set("userId", claims.ID)
// 		c.Next()
// 	}
// }

var (
	ErrTokenExpired     = errors.New("token is expired")
	ErrTokenNotValidYet = errors.New("token not active yet")
	ErrTokenMalformed   = errors.New("that's not even a token")
	ErrTokenInvalid     = errors.New("couldn't handle this token")
)

// 创建一个token
func (j *JWT) CreateToken(claims models.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*models.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, ErrTokenInvalid

	} else {
		return nil, ErrTokenInvalid

	}

}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", ErrTokenInvalid
}
