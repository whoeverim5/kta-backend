// Package utils /*
/**
 * @Author:      吴贤权
 * @Date:        2022/10/3 下午4:48
 * @Description: 获取随即验证码工具
**/
package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GetCode(width int) (string, error) {
	var code strings.Builder
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rl := len(numeric)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < width; i++ {
		if _, err := fmt.Fprintf(&code, "%d", numeric[r.Intn(rl)]); err != nil {
			return "", err
		}
	}
	return code.String(), nil
}
