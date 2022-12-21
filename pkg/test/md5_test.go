// Package test /*
/**
 * @Author:      吴贤权
 * @Date:        2022/10/3 下午4:47
 * @Description: 测试md5加密
**/
package test

import (
	"fmt"
	"log"
	"testing"

	"ktaexam/pkg/utils"
)

func TestMd5(t *testing.T) {
	rSalt := "iuaaunWaexqi-huz70%au-ix8a-ua-1a"
	md5Pwd := utils.Md5Crypto("f7a71f7bdb10f7e78a5ae09ad74054cc", rSalt)
	log.Println(md5Pwd)
	// daa1dc0e39bc43908b7445940f2d3fda
}

func TestRandomSalt(t *testing.T) {
	s := utils.GetRandomSalt()
	fmt.Println(s)
}
