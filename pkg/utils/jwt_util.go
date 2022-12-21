// Package utils /*
/**
 * @Author:      吴贤权
 * @Date:        2022/10/3 下午4:48
 * @Description: 用户鉴权工具
**/
package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWT 签名结构
type JWT struct {
	Secret []byte
}

// 一些常量
var (
	TokenExpired           = errors.New("token is expired")
	TokenNotValidYet       = errors.New("token not active yet")
	TokenMalformed         = errors.New("that's not even a token")
	TokenInvalid           = errors.New("couldn't handle this token")
	Secret                 = "exam-wu-key-secret"
	TokenExpireTime  int64 = 1 * 24 * 60 * 60
)

// CustomClaims 载荷
type CustomClaims struct {
	Account string `json:"account"`
	UserId  string `json:"user_id"`
	jwt.StandardClaims
}

// NewJWT 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSecret()),
	}
}

// GetSecret 获取signKey
func GetSecret() string {
	return Secret
}

// SetSecret 修改SignKey
func SetSecret(key string) string {
	Secret = key
	return Secret
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.Secret)
}

// ParseToken 解析Token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.Secret, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// RefreshToken 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.Secret, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
