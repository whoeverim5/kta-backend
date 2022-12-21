// Package test /*
/**
 * @Author:      吴贤权
 * @Date:        2022/10/3 下午4:47
 * @Description: 测试uuid生成
**/
package test

import (
	"log"
	"testing"

	"ktaexam/pkg/utils"
)

func TestNewUserId(t *testing.T) {
	userId := utils.NewUserId()
	log.Println(userId)
}
