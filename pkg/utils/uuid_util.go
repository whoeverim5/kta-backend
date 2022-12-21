// Package utils /*
/**
 * @Author:      吴贤权
 * @Date:        2022/10/3 下午4:49
 * @Description: 生成uuid工具
**/
package utils

import "github.com/google/uuid"

func NewUserId() string {
	return uuid.NewString()
}
