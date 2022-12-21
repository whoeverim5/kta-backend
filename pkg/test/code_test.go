// Package test /*
/**
 * @Author:      吴贤权
 * @Date:        2022/10/3 下午4:46
 * @Description: 测试获取随即验证码
**/
package test

import (
	"fmt"
	"testing"

	"ktaexam/pkg/utils"
)

func TestCode(t *testing.T) {
	code, _ := utils.GetCode(6)
	fmt.Println(code)
}
