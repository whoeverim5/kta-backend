// Package utils /*
/**
 * @Author:      吴贤权
 * @Date:        2022/10/3 下午4:48
 * @Description: 密码md5加密工具
**/
package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

var (
	salt = "exam-zucchini-WuAquarian%5250178" // md5加密盐
)

func Md5Crypto(str, salt string) string {
	bSalt := []byte(salt)
	bStr := []byte(str)
	crypto := md5.New()
	crypto.Write(bSalt)
	crypto.Write(bStr)
	return hex.EncodeToString(crypto.Sum(nil))
}

func GetRandomSalt() string {
	randomSalt := ""
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len(salt); i++ {
		randomSalt += string(salt[r.Intn(len(salt))])
	}
	return randomSalt
}
