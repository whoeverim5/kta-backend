// Package test /*
/**
 * @Author:      吴贤权
 * @Date:        2022/10/3 下午4:47
 * @Description: 测试token生成
**/
package test

import (
	"log"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"ktaexam/pkg/utils"
)

func TestToken(t *testing.T) {
	claims := utils.CustomClaims{
		Account: "1192445074",
		UserId:  "user_test",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(utils.TokenExpireTime)).Unix(),
			Issuer:    "wu-" + time.Now().Format("06-01-02")},
	}
	token, _ := utils.NewJWT().CreateToken(claims)
	log.Println("token: ", token)
	c, _ := utils.NewJWT().ParseToken(token)
	log.Println("account:", c.Account)
	log.Println("user_id:", c.UserId)
	log.Println("Expire:", c.ExpiresAt)
	log.Println("Issuer:", c.Issuer)
}
