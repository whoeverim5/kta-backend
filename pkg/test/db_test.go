// Package test /*
/**
 * @Author:      吴贤权
 * @Date:        2022/10/3 下午4:46
 * @Description: 测试mysql数据库连接
**/
package test

import (
	"fmt"
	"testing"

	"ktaexam/pkg/dao"
	"ktaexam/pkg/model"
)

func TestDB(t *testing.T) {
	users := make([]model.User, 0)
	dao.DB.Find(&users)
	fmt.Println(users)
}
